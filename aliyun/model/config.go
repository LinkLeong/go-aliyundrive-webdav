package model

const (
	APIBASE            = "https://api.aliyundrive.com"
	APILISTURL         = APIBASE + "/adrive/v3/file/list"
	APIREFRESHTOKENURL = APIBASE + "/token/refresh"
)

type Config struct {
	RefreshToken string `json:"refresh_token"`
	Token        string `json:"token"`
	DriveId      string `json:"drive_id"`
	ExpireTime   int64  `json:"expire_time"`
}
