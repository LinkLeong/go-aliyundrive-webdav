package aliyun

import (
	"encoding/json"
	"fmt"
	"go-aliyun/aliyun/cache"
	"go-aliyun/aliyun/model"
	"go-aliyun/aliyun/net"
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
	postData["limit"] = 100
	postData["all"] = true
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

func GetFile(url string, token string) []byte {

	body := net.Get(url, token)

	return body
}

func RefreshToken(refreshToken string) model.RefreshTokenModel {
	rs := net.Post(model.APIREFRESHTOKENURL, "", []byte(`{"refresh_token":"`+refreshToken+`"}`))
	var refresh model.RefreshTokenModel
	if len(rs) > 0 {
		err := json.Unmarshal(rs, &refresh)
		if err != nil {
			fmt.Println("刷新token失败,失败信息", err)
			fmt.Println("刷新token返回信息", refresh)
		}
	} else {
		fmt.Println("刷新token失败")
	}
	return refresh

}
