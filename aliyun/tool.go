package aliyun

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
)

//处理内容
func ContentHandle(r *http.Request, token string, driveId string, parentId string, fileName string) {
	//需要判断参数里面的有效期
	//默认截取长度10485760
	//const DEFAULT int64 = 10485760
	const DEFAULT int64 = 1048576
	var count float64 = 1
	var total int64 = 0
	byteSize := DEFAULT

	if len(parentId) == 0 {
		parentId = "root"
	}
	count = math.Ceil(float64(r.ContentLength) / float64(DEFAULT))
	uploadUrl, uploadId, fileId := UpdateFileFile(token, driveId, fileName, parentId, strconv.FormatInt(r.ContentLength, 10), int(count))
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
		n, err := io.ReadFull(r.Body, dataByte)
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
		UploadFile(uploadUrl[i].Str, token, dataByte)
	}

	UploadFileComplete(token, driveId, uploadId, fileId, parentId)
}
