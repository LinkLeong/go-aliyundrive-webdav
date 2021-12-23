package model

const (
	APIBASE            = "https://api.aliyundrive.com"
	APILISTURL         = APIBASE + "/adrive/v3/file/list"
	APIFILEPATH        = APIBASE + "/adrive/v1/file/get_path"
	APIREFRESHTOKENURL = APIBASE + "/token/refresh"
	APIREMOVETRASH     = APIBASE + "/v2/recyclebin/trash" //移动到垃圾箱
	APIFILEUPDATE      = APIBASE + "/v3/file/update"
	APIMKDIR           = APIBASE + "/adrive/v2/file/createWithFolders"
	APIFILEDETAIL      = APIBASE + "/v2/file/get"
	APIFILEBATCH       = APIBASE + "/v3/batch"
	APIFILEUPLOAD      = APIBASE + "/adrive/v2/file/createWithFolders"
	APIFILEUPLOADURL   = APIBASE + "/v2/file/get_upload_url"
	APIFILEUPLOADFILE  = APIBASE + "/v2/file/create_with_proof" //"/v2/file/create"
	APIFILECOMPLETE    = APIBASE + "/v2/file/complete"
	APIFILEDOWNLOAD    = APIBASE + "/v2/file/get_download_url"
	APITOTLESIZE       = APIBASE + "/v2/databox/get_personal_info"
	APISEARCH          = APIBASE + "/adrive/v3/file/search"
)

type Config struct {
	RefreshToken string `json:"refresh_token"`
	Token        string `json:"token"`
	DriveId      string `json:"drive_id"`
	ExpireTime   int64  `json:"expire_time"`
}
