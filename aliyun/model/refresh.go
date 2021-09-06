package model

type RefreshTokenModel struct {
	AccessToken    string `json:"access_token"`
	DefaultDriveId string `json:"default_drive_id"`
	RefreshToken   string `json:"refresh_token"`
	ExpiresIn      int64  `json:"expires_in"`
}
