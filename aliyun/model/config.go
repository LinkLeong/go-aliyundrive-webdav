package model

const (
	APIBASE            = "https://api.aliyundrive.com"
	APILISTURL         = APIBASE + "/adrive/v3/file/list"
	APIREFRESHTOKENURL = APIBASE + "/token/refresh"
	APIREMOVETRASH     = APIBASE + "/v2/recyclebin/trash" //移动到垃圾箱
	APIFILEUPDATE      = APIBASE + "/v3/file/update"
	APIMKDIR           = APIBASE + "/adrive/v2/file/createWithFolders"
)

type Config struct {
	RefreshToken string `json:"refresh_token"`
	Token        string `json:"token"`
	DriveId      string `json:"drive_id"`
	ExpireTime   int64  `json:"expire_time"`
}
