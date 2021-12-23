package aliyun

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"go-aliyun-webdav/aliyun/cache"
	"go-aliyun-webdav/aliyun/model"
	"go-aliyun-webdav/aliyun/net"
	"go-aliyun-webdav/utils"
	"io"
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
	var total int64 = 0
	byteSize := DEFAULT

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
	buff, _ := ioutil.ReadAll(r.Body)
	preHashReader := ioutil.NopCloser(bytes.NewReader(buff))
	fullHashReader := ioutil.NopCloser(bytes.NewReader(buff))
	partsReader := ioutil.NopCloser(bytes.NewReader(buff))
	proofReader := ioutil.NopCloser(bytes.NewReader(buff))

	prehashdataByte := make([]byte, 1024)
	_, err := io.ReadFull(preHashReader, prehashdataByte)
	if err != nil {
		fmt.Println("IO Error", fileName)
	}
	h := sha1.New()
	h.Write(prehashdataByte)

	//检查是否可以极速上传，逻辑如下
	//取文件的前1K字节，做SHA1摘要，调用创建文件接口，pre_hash参数为SHA1摘要，如果返回409，则这个文件可以极速上传
	prehashData := `{"drive_id":"` + driveId + `","parent_file_id":"` + parentId + `","name":"` + fileName + `","type":"file","check_name_mode":"overwrite","size":` + strconv.FormatInt(r.ContentLength, 10) + `,"pre_hash":"` + hex.EncodeToString(h.Sum(nil)) + `","proof_version":"v1"}`
	_, code := net.PostExpectStatus(model.APIFILEUPLOAD, token, []byte(prehashData))
	var offset int64 = 0
	var proof string = ""
	var flashUpload bool = false
	if code == 409 {
		md := md5.New()
		tokenBytes, _ := ioutil.ReadAll(strings.NewReader(token))
		md.Write(tokenBytes)
		tokenMd5 := hex.EncodeToString(md.Sum(nil))
		tokenMd5 = tokenMd5[:16]
		f, _ := strconv.ParseInt(tokenMd5, 16, 64)
		offset = f % r.ContentLength
		//offset2, _ := utils.CalcOffSet(token, strconv.FormatInt(r.ContentLength, 10))
		//fmt.Println(offset2)
		offsetBytes := make([]byte, 8)
		all, _ := io.ReadAll(proofReader)
		_, err := bytes.NewReader(all).ReadAt(offsetBytes, offset)
		if err == nil {
			proof = utils.GetProof(offsetBytes)
		}
		flashUpload = true
	}
	h2 := sha1.New()
	all, _ := io.ReadAll(fullHashReader)
	h2.Write(all)

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
		if r.ContentLength-total > DEFAULT {
			byteSize = DEFAULT
		} else {
			byteSize = r.ContentLength - total
		}
		dataByte := make([]byte, byteSize)
		n, err := io.ReadFull(partsReader, dataByte)
		//n, err := r.Body.Read(dataByte)
		if err != nil {
			fmt.Println("获取字节内容出错", err)
		}
		total += int64(n)
		//	fmt.Println("对比长度", total)
		//	fmt.Println("提交数据的长度", len(dataByte))

		//	u, _ := url.Parse(uploadUrl[i].Str)
		//	params := u.Query()
		//	fmt.Println(params.Get("x-oss-expires"))
		//fmt.Println(uploadUrl[i])
		UploadFile(uploadUrl[i].Str, token, dataByte)
	}

	UploadFileComplete(token, driveId, uploadId, fileId, parentId)
	cache.GoCache.Delete(parentId)
	return fileId
}
