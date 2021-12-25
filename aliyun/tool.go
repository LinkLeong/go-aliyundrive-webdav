package aliyun

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/tidwall/gjson"
	"go-aliyun-webdav/aliyun/cache"
	"go-aliyun-webdav/aliyun/model"
	"go-aliyun-webdav/aliyun/net"
	"go-aliyun-webdav/utils"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//处理内容
func ContentHandle(r *http.Request, token string, driveId string, parentId string, fileName string) string {
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
		return ""
	}

	//proof 偏移量
	var offset int64 = 0
	//proof内容base64
	var proof string = ""
	//是否闪传
	var flashUpload bool = false
	//status code
	var code int
	var readbytes []byte
	var uploadUrl []gjson.Result
	var uploadId string
	var uploadFileId string
	//大于150K小于1G的才开启闪传，文件太大会导致内存溢出
	//由于webdav协议的局限性，无法预先计算文件校验hash
	if r.ContentLength > 1024*150 && r.ContentLength <= 1024*1024*1024 {
		preHashDataBytes := make([]byte, 1024)
		_, err := io.ReadFull(r.Body, preHashDataBytes)
		if err != nil {
			fmt.Println("error reading file", fileName)
			return ""
		}
		readbytes = append(readbytes, preHashDataBytes...)
		h := sha1.New()
		h.Write(preHashDataBytes)
		//检查是否可以极速上传，逻辑如下
		//取文件的前1K字节，做SHA1摘要，调用创建文件接口，pre_hash参数为SHA1摘要，如果返回409，则这个文件可以极速上传
		preHashRequest := `{"drive_id":"` + driveId + `","parent_file_id":"` + parentId + `","name":"` + fileName + `","type":"file","check_name_mode":"overwrite","size":` + strconv.FormatInt(r.ContentLength, 10) + `,"pre_hash":"` + hex.EncodeToString(h.Sum(nil)) + `","proof_version":"v1"}`
		_, code = net.PostExpectStatus(model.APIFILEUPLOAD, token, []byte(preHashRequest))
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
			var offsetBytes []byte
			if end < 1024 {
				offsetBytes = readbytes[offset:int64(end)]
				proof = utils.GetProof(offsetBytes)
			} else {
				//先读取到offset end位置的所有字节，由于上面已经读取1024，这里剪掉
				offsetBytes = make([]byte, int64(end-1024))
				_, err2 := io.ReadFull(r.Body, offsetBytes)
				if err2 != nil {
					fmt.Println(err2)
					return ""
				}
				readbytes = append(readbytes, offsetBytes...)
				offsetBytes = offsetBytes[offset-1024 : int64(end)-1024]
				proof = utils.GetProof(offsetBytes)
			}
			flashUpload = true
		}
		buff := make([]byte, r.ContentLength-int64(len(readbytes)))
		_, err3 := io.ReadFull(r.Body, buff)
		if err3 != nil {
			fmt.Println(err3)
			return ""
		}
		h2 := sha1.New()
		readbytes = append(readbytes, buff...)
		h2.Write(readbytes)
		uploadUrl, uploadId, uploadFileId = UpdateFileFile(token, driveId, fileName, parentId, strconv.FormatInt(r.ContentLength, 10), int(count), strings.ToUpper(hex.EncodeToString(h2.Sum(nil))), proof, flashUpload)
	} else {
		uploadUrl, uploadId, uploadFileId = UpdateFileFile(token, driveId, fileName, parentId, strconv.FormatInt(r.ContentLength, 10), int(count), "", "", false)
	}

	if flashUpload && (uploadFileId != "") {
		//UploadFileComplete(token, driveId, uploadId, uploadFileId, parentId)
		cache.GoCache.Delete(parentId)
		return uploadFileId
	}
	if len(uploadUrl) == 0 {
		return ""
	}
	var bg time.Time = time.Now()
	for i := 0; i < int(count); i++ {
		var dataByte []byte
		if int(count) == 1 {
			dataByte = make([]byte, r.ContentLength)
		} else if i == int(count)-1 {
			dataByte = make([]byte, r.ContentLength-int64(i)*DEFAULT)
		} else {
			dataByte = make([]byte, DEFAULT)
		}
		_, err := io.ReadFull(r.Body, dataByte)
		if err != nil {
			fmt.Println("err reading from request body", err, fileName, uploadId)
			_, err := io.ReadFull(bytes.NewReader(readbytes), dataByte)
			if err != nil {
				return ""
			}
		}
		UploadFile(uploadUrl[i].Str, token, dataByte)
		fmt.Println("multithread upload, thread #", i)

	}
	fmt.Println("uploading done, elapsed ", time.Now().Sub(bg).String())
	UploadFileComplete(token, driveId, uploadId, uploadFileId, parentId)
	cache.GoCache.Delete(parentId)
	return uploadFileId
}
