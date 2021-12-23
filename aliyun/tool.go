package aliyun

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"go-aliyun-webdav/aliyun/cache"
	"go-aliyun-webdav/aliyun/model"
	"go-aliyun-webdav/aliyun/net"
	"go-aliyun-webdav/utils"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
)

//处理内容
func ContentHandle(r *http.Request, token string, driveId string, parentId string, fileName string) (fileId string) {
	//需要判断参数里面的有效期
	//默认截取长度10485760
	//const DEFAULT int64 = 10485760
	const DEFAULT int64 = 10485760
	var count float64 = 1

	if len(parentId) == 0 {
		parentId = "root"
	}
	if r.ContentLength > 0 {
		count = math.Ceil(float64(r.ContentLength) / float64(DEFAULT))
	} else {
		//dataTemp, _ := io.ReadAll(r.Body)
		//r.ContentLength = int64(len(dataTemp))
		return
	}
	//proof 偏移量
	var offset int64 = 0
	//proof内容base64
	var proof string = ""
	//是否闪传
	var flashUpload bool = false
	//status code
	var code int
	//获取请求文件内容，对于大文件不是很友好，但是为了算文件hash好像没有什么更好的方法
	buff, _ := ioutil.ReadAll(r.Body)
	if r.ContentLength > 1024*1024*20 {
		preHashDataBytes := buff[:1024]
		h := sha1.New()
		h.Write(preHashDataBytes)
		//检查是否可以极速上传，逻辑如下
		//取文件的前1K字节，做SHA1摘要，调用创建文件接口，pre_hash参数为SHA1摘要，如果返回409，则这个文件可以极速上传
		preHashRequest := `{"drive_id":"` + driveId + `","parent_file_id":"` + parentId + `","name":"` + fileName + `","type":"file","check_name_mode":"overwrite","size":` + strconv.FormatInt(r.ContentLength, 10) + `,"pre_hash":"` + hex.EncodeToString(h.Sum(nil)) + `","proof_version":"v1"}`
		_, code = net.PostExpectStatus(model.APIFILEUPLOAD, token, []byte(preHashRequest))
	}
	if code == 409 {
		md := md5.New()
		tokenBytes := []byte(token)
		md.Write(tokenBytes)
		tokenMd5 := hex.EncodeToString(md.Sum(nil))
		first16 := tokenMd5[:16]
		f, err := strconv.ParseUint(first16, 16, 64)
		if err != nil {
			fmt.Println(err)
		}
		offset = int64(f % uint64(r.ContentLength))
		end := math.Min(float64(offset+8), float64(r.ContentLength))
		offsetBytes := buff[offset:int64(end)]
		proof = utils.GetProof(offsetBytes)
		flashUpload = true
	}
	h2 := sha1.New()
	h2.Write(buff)
	uploadUrl, uploadId, fileId := UpdateFileFile(token, driveId, fileName, parentId, strconv.FormatInt(r.ContentLength, 10), int(count), strings.ToUpper(hex.EncodeToString(h2.Sum(nil))), proof, flashUpload)

	if flashUpload && (fileId != "") {
		UploadFileComplete(token, driveId, uploadId, fileId, parentId)
		cache.GoCache.Delete(parentId)
		return fileId
	}
	if len(uploadUrl) == 0 {
		return
	}
	for i := 0; i < int(count); i++ {

		var dataByte []byte
		if i == int(count)-1 {
			dataByte = buff[int64(i)*DEFAULT : r.ContentLength]
		} else {
			dataByte = buff[int64(i)*DEFAULT : int64(i+1)*DEFAULT]
		}
		UploadFile(uploadUrl[i].Str, token, dataByte)
	}

	UploadFileComplete(token, driveId, uploadId, fileId, parentId)
	cache.GoCache.Delete(parentId)
	return fileId
}
