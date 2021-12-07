package aliyun

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-aliyun-webdav/aliyun/cache"
	"go-aliyun-webdav/aliyun/model"
	"go-aliyun-webdav/aliyun/net"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/tidwall/gjson"
)

func GetList(token string, driveId string, parentFileId string) (model.FileListModel, error) {

	if len(parentFileId) == 0 {
		parentFileId = "root"
	}

	var list model.FileListModel
	if result, ok := cache.GoCache.Get(parentFileId); ok {
		list, ok = result.(model.FileListModel)
		if ok {
			return list, nil
		}
	}

	postData := make(map[string]interface{})
	postData["drive_id"] = driveId
	postData["parent_file_id"] = parentFileId
	postData["limit"] = 200
	postData["all"] = false
	postData["url_expire_sec"] = 1600
	postData["image_thumbnail_process"] = "image/resize,w_400/format,jpeg"
	postData["image_url_process"] = "image/resize,w_1920/format,jpeg"
	postData["video_thumbnail_process"] = "video/snapshot,t_0,f_jpg,ar_auto,w_300"
	postData["fields"] = "*"
	postData["order_by"] = "updated_at"
	postData["order_direction"] = "DESC"

	data, err := json.Marshal(postData)
	if err != nil {
		fmt.Println("获取列表转义数据失败", err)
		return model.FileListModel{}, err
	}

	body := net.Post(model.APILISTURL, token, data)

	e := json.Unmarshal(body, &list)
	if e != nil {
		fmt.Println(e)
	}
	if len(list.Items) > 0 {
		cache.GoCache.SetDefault(parentFileId, list)
	}
	return list, nil
}

func GetFilePath(token string, driveId string, parentFileId string, fileId string, typeStr string) (string, error) {

	if len(parentFileId) == 0 {
		parentFileId = "root"
	}
	path := "/"
	var list model.ListFilePath
	if result, ok := cache.GoCache.Get(parentFileId + "path"); ok {
		path, ok = result.(string)
		if ok {
			return path, nil
		}
	}

	postData := make(map[string]interface{})
	postData["drive_id"] = driveId
	postData["file_id"] = fileId

	data, err := json.Marshal(postData)
	if err != nil {
		fmt.Println("获取列表转义数据失败", err)
		return "/", err
	}

	body := net.Post(model.APIFILEPATH, token, data)

	e := json.Unmarshal(body, &list)
	if e != nil {
		fmt.Println(e)
	}
	minNum := 0
	if typeStr == "folder" {
		minNum = 1
	}
	for i := len(list.Items); i > minNum; i-- {
		if list.Items[i-1].Type == "folder" {
			path += list.Items[i-1].Name + "/"
		}
	}

	cache.GoCache.SetDefault(parentFileId+"path", path)

	return path, nil
}

func GetFile(w http.ResponseWriter, url string, token string, rangeStr string, ifRange string) bool {

	body := net.Get(w, url, token, rangeStr, ifRange)
	//net.GetProxy(w, req, url, token)
	return body
	//return []byte{}
}

func RefreshToken(refreshToken string) model.RefreshTokenModel {
	path := refreshToken
	if _, errs := os.Stat(path); errs == nil {
		buf, _ := ioutil.ReadFile(path)
		refreshToken = string(buf)
		if len(refreshToken) >= 32 {
			refreshToken = refreshToken[:32] // refreshToken is only 32 bit?? FIXME
		}
	}
	rs := net.Post(model.APIREFRESHTOKENURL, "", []byte(`{"refresh_token":"`+refreshToken+`"}`))
	var refresh model.RefreshTokenModel

	if len(rs) <= 0 {
		fmt.Println("刷新token失败")
		return refresh
	}

	err := json.Unmarshal(rs, &refresh)
	if err != nil {
		fmt.Println("刷新token失败,失败信息", err)
		fmt.Println("刷新token返回信息", refresh)
		return refresh
	}

	if refreshToken == refresh.RefreshToken {
		return refresh
	}

	_, err = os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return refresh
	}
	if err != nil {
		fmt.Println("更新token文件失败,失败信息", err)
		return refresh
	}

	err = ioutil.WriteFile(path, []byte(refresh.RefreshToken), 0600)
	if err != nil {
		fmt.Println("更新token文件失败,失败信息", err)
	}

	return refresh
}

func RemoveTrash(token string, driveId string, fileId string, parentFileId string) bool {
	rs := net.Post(model.APIREMOVETRASH, token, []byte(`{"drive_id":"`+driveId+`","file_id":"`+fileId+`"}`))
	if len(rs) == 0 {
		cache.GoCache.Delete(parentFileId)
	}
	return false
}

func ReName(token string, driveId string, newName string, fileId string) bool {
	rs := net.Post(model.APIFILEUPDATE, token, []byte(`{"drive_id":"`+driveId+`","file_id":"`+fileId+`","name":"`+newName+`","check_name_mode":"refuse"}`))
	var m model.ListModel
	e := json.Unmarshal(rs, &m)
	if e != nil {
		fmt.Println(e)
	}
	cache.GoCache.Delete(m.ParentFileId)
	fmt.Println(string(rs))
	return true
}
func MakeDir(token string, driveId string, name string, parentFileId string) bool {
	rs := net.Post(model.APIMKDIR, token, []byte(`{"drive_id":"`+driveId+`","parent_file_id":"`+parentFileId+`","name":"`+name+`","check_name_mode":"refuse","type":"folder"}`))
	//正确返回示例
	//{
	//	"parent_file_id": "root",
	//	"type": "folder",
	//	"file_id": "6134d1b4253b74c8f7e24d72afa20f58fd19ac28",
	//	"domain_id": "bj29",
	//	"drive_id": "1662258",
	//	"file_name": "新0000",
	//	"encrypt_mode": "none"
	//}
	if gjson.GetBytes(rs, "file_name").Str == name {
		cache.GoCache.Delete(parentFileId)
	}
	return true
}

func GetFileDetail(token string, driveId string, fileId string) model.ListModel {
	rs := net.Post(model.APIFILEDETAIL, token, []byte(`{"drive_id":"`+driveId+`","file_id":"`+fileId+`"}`))
	var m model.ListModel
	e := json.Unmarshal(rs, &m)
	if e != nil {
		fmt.Println(e)
	}
	return m
}

func BatchFile(token string, driveId string, fileId string, parentFileId string) bool {

	//	{
	//		"requests": ,
	//	"resource": "file"
	//	}

	var bodyJson string = `{"drive_id": "` + driveId + `","file_id": "` + fileId + `","to_drive_id": "` + driveId + `","to_parent_file_id": "` + parentFileId + `"}`
	var contentType string = `{"Content-Type": "application/json"}`

	var requests string = `{"requests":[{"body": ` + bodyJson + `,"headers": ` + contentType + `,"id": "` + fileId + `","method": "POST","url": "/file/move"}],"resource": "file"}`

	rs := net.Post(model.APIFILEBATCH, token, []byte(requests))
	if gjson.GetBytes(rs, "responses.0.friends").Num == 200 {
		cache.GoCache.Delete(parentFileId)
		cache.GoCache.Delete(fileId)
		return true
	}

	return false
}
func UpdateFileFolder(token string, driveId string, fileName string, parentFileId string) bool {

	//	{
	//		"requests": ,
	//	"resource": "file"
	//	}
	createData := `{"drive_id": "` + driveId + `","parent_file_id": "` + parentFileId + `","name": "` + fileName + `","check_name_mode": "refuse","type": "folder"}`
	net.Post(model.APIFILEUPLOAD, token, []byte(createData))
	// rs := net.Post(model.APIFILEUPLOAD, token, []byte(createData))
	// fmt.Println(string(rs))
	//正确返回占星显示
	//	{"parent_file_id":"60794ad941ee2d8d24f843b7a0ffd80279927dfc","type":"folder","file_id":"613caeb4d5b1ba9fb4604d4aa5aef2b408ab3121","domain_id":"bj29","drive_id":"1662258","file_name":"1SDSDSD.png","encrypt_mode":"none"}
	//
	//	{
	//		"parent_file_id": "root",
	//		"part_info_list": [
	//	{
	//		"part_number": 1,
	//		"upload_url": "https://bj29.cn-beijing.data.alicloudccp.com/igQPcuUn%2F1662258%2F613a1091919bb599f4ac4917bfe16af6b7066795%2F613a10919ab3804e88e846ee9ea459de51d8f58d?partNumber=1&uploadId=BD8449BB161A4F54A1252E3B5B121641&x-oss-access-key-id=LTAIsE5mAn2F493Q&x-oss-expires=1631198881&x-oss-signature=wp2WCgyfqxZhJH%2BsPaw6XASRKXHa92p3e9NOjcN4Ui8%3D&x-oss-signature-version=OSS2",
	//		"internal_upload_url": "http://ccp-bj29-bj-1592982087.oss-cn-beijing-internal.aliyuncs.com/igQPcuUn%2F1662258%2F613a1091919bb599f4ac4917bfe16af6b7066795%2F613a10919ab3804e88e846ee9ea459de51d8f58d?partNumber=1&uploadId=BD8449BB161A4F54A1252E3B5B121641&x-oss-access-key-id=LTAIsE5mAn2F493Q&x-oss-expires=1631198881&x-oss-signature=wp2WCgyfqxZhJH%2BsPaw6XASRKXHa92p3e9NOjcN4Ui8%3D&x-oss-signature-version=OSS2",
	//		"content_type": ""
	//	}
	//],
	//	"upload_id": "BD8449BB161A4F54A1252E3B5B121641",
	//	"rapid_upload": false,
	//	"type": "file",
	//	"file_id": "613a1091919bb599f4ac4917bfe16af6b7066795",
	//	"domain_id": "bj29",
	//	"drive_id": "1662258",
	//	"file_name": "photo_1614943806132229.jpg",
	//	"encrypt_mode": "none",
	//	"location": "cn-beijing"
	//	}

	return false
}

func UpdateFileFile(token string, driveId string, fileName string, parentFileId string, size string, length int) ([]gjson.Result, string, string) {

	if len(parentFileId) == 0 {
		parentFileId = "root"
	}

	var partStr string = "["
	for i := 0; i < length; i++ {
		partStr += `{"part_number":` + strconv.Itoa(i+1) + `},`
	}
	partStr = partStr[:len(partStr)-1]
	partStr += "]"
	createData := `{"drive_id":"` + driveId + `","part_info_list":` + partStr + `,"parent_file_id":"` + parentFileId + `","name":"` + fileName + `","type":"file","check_name_mode":"auto_rename","size":` + size + `,"content_hash_name":"none","proof_version":"v1"}`
	rs := net.Post(model.APIFILEUPLOADFILE, token, []byte(createData))
	urlArr := gjson.GetBytes(rs, "part_info_list.#.upload_url").Array()
	if len(urlArr) == 0 {
		fmt.Println("创建文件出错", string(rs))
	}
	return urlArr, gjson.GetBytes(rs, "upload_id").Str, gjson.GetBytes(rs, "file_id").Str
	//正确返回占星显示
	//
	//	{
	//		"parent_file_id": "root",
	//		"part_info_list": [
	//	{
	//		"part_number": 1,
	//		"upload_url": "https://bj29.cn-beijing.data.alicloudccp.com/igQPcuUn%2F1662258%2F613a1091919bb599f4ac4917bfe16af6b7066795%2F613a10919ab3804e88e846ee9ea459de51d8f58d?partNumber=1&uploadId=BD8449BB161A4F54A1252E3B5B121641&x-oss-access-key-id=LTAIsE5mAn2F493Q&x-oss-expires=1631198881&x-oss-signature=wp2WCgyfqxZhJH%2BsPaw6XASRKXHa92p3e9NOjcN4Ui8%3D&x-oss-signature-version=OSS2",
	//		"internal_upload_url": "http://ccp-bj29-bj-1592982087.oss-cn-beijing-internal.aliyuncs.com/igQPcuUn%2F1662258%2F613a1091919bb599f4ac4917bfe16af6b7066795%2F613a10919ab3804e88e846ee9ea459de51d8f58d?partNumber=1&uploadId=BD8449BB161A4F54A1252E3B5B121641&x-oss-access-key-id=LTAIsE5mAn2F493Q&x-oss-expires=1631198881&x-oss-signature=wp2WCgyfqxZhJH%2BsPaw6XASRKXHa92p3e9NOjcN4Ui8%3D&x-oss-signature-version=OSS2",
	//		"content_type": ""
	//	}
	//],
	//	"upload_id": "BD8449BB161A4F54A1252E3B5B121641",
	//	"rapid_upload": false,
	//	"type": "file",
	//	"file_id": "613a1091919bb599f4ac4917bfe16af6b7066795",
	//	"domain_id": "bj29",
	//	"drive_id": "1662258",
	//	"file_name": "photo_1614943806132229.jpg",
	//	"encrypt_mode": "none",
	//	"location": "cn-beijing"
	//	}

	//return false
}
func UploadFile(url string, token string, data []byte) {
	net.Put(url, token, data)
}
func UploadFileComplete(token string, driveId string, uploadId string, fileId string, parentId string) bool {
	//	private String drive_id;
	//	private String file_id;
	//	private String upload_id;
	//	{
	//		"requests": ,
	//	"resource": "file"
	//	}
	createData := `{"drive_id": "` + driveId + `","file_id": "` + fileId + `","upload_id":"` + uploadId + `"}`

	rs := net.Post(model.APIFILECOMPLETE, token, []byte(createData))
	fmt.Println(string(rs))
	//正确返回占星显示
	//	}
	cache.GoCache.Delete(parentId)

	return false
}
func GetDownloadUrl(token string, driveId string, fileId string) string {

	postData := make(map[string]interface{})
	postData["drive_id"] = driveId
	postData["file_id"] = fileId

	data, _ := json.Marshal(postData)

	body := net.Post(model.APIFILEDOWNLOAD, token, data)
	return gjson.GetBytes(body, "url").Str

}
func GetBoxSize(token string) (string, string) {

	postData := make(map[string]interface{})

	data, _ := json.Marshal(postData)

	body := net.Post(model.APITOTLESIZE, token, data)
	return gjson.GetBytes(body, "personal_space_info.total_size").String(), gjson.GetBytes(body, "personal_space_info.used_size").String()

}
