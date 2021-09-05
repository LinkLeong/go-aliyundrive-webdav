package model

import "time"

type ListModel struct {
	DriveId       string    `json:"drive_id"`
	FileId        string    `json:"file_id"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	Status        string    `json:"status"`
	ParentFileId  string    `json:"parent_file_id"`
	Starred       bool      `json:"starred"`
	ContentType   string    `json:"content_type"`
	FileExtension string    `json:"file_extension"`
	MimeType      string    `json:"mime_type"`
	MimeExtension string    `json:"mime_extension"`
	Hidden        bool      `json:"hidden"`
	Size          int       `json:"size"`
	Category      string    `json:"category"`
	DownloadUrl   string    `json:"download_url"`
	Url           string    `json:"url"`
	Thumbnail     string    `json:"thumbnail"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type FileListModel struct {
	Items      []ListModel `json:"items"`
	NextMarker string      `json:"next_marker"`
}

//{
//"drive_id": "1662258",
//"domain_id": "bj29",
//"file_id": "605df29f93398d72d9cf487eba5f34ae3fb61637",
//"name": "photo_1615568241006.jpg",
//"type": "file",
//"content_type": "image/jpeg",
//"created_at": "2021-03-26T14:41:35.973Z",
//"updated_at": "2021-03-26T14:41:43.141Z",
//"file_extension": "jpg",
//"mime_type": "image/jpeg",
//"mime_extension": "jpg",
//"hidden": false,
//"size": 826909,
//"starred": false,
//"status": "available",
//"user_meta": "{\"size\":826909,\"android_local_file_path\":\"/storage/emulated/0/Lark/camera/photo/photo_1615568241006.jpg\",\"android_identify_id\":\"258092\",\"time\":1615568241000,\"hash\":\"CED19A2206F6C30E408D3666D9875D10CF7B615D\"}",
//"labels": [
//"海报",
//"小册子",
//"文本",
//"手机截图"
//],
//"upload_id": "04D1A860D76D45A382E81D5D758CE577",
//"parent_file_id": "root",
//"crc64_hash": "10067593862750608359",
//"content_hash": "E87427276113C30290DC551D81EDDB59EAC271AD",
//"content_hash_name": "sha1",
//"download_url": "https://bj29.cn-beijing.data.alicloudccp.com/f9n9Bwcf%2F1662258%2F605df29f93398d72d9cf487eba5f34ae3fb61637%2F605df29f031042c8b1714539afa2aa7b4e28feb5?di=bj29&dr=1662258&f=605df29f93398d72d9cf487eba5f34ae3fb61637&response-content-disposition=attachment%3B%20filename%2A%3DUTF-8%27%27photo_1615568241006.jpg&u=ff08e5f172214ee88eb7790594e206b5&x-oss-access-key-id=LTAIsE5mAn2F493Q&x-oss-additional-headers=referer&x-oss-expires=1630648496&x-oss-signature=VOlfM4za0Do9GQ%2FRD48loINNcGLSm3wWX51qs1021Jc%3D&x-oss-signature-version=OSS2",
//"url": "https://bj29.cn-beijing.data.alicloudccp.com/f9n9Bwcf%2F1662258%2F605df29f93398d72d9cf487eba5f34ae3fb61637%2F605df29f031042c8b1714539afa2aa7b4e28feb5?x-oss-access-key-id=LTAIsE5mAn2F493Q&x-oss-additional-headers=referer&x-oss-expires=1630648496&x-oss-process=image%2Fresize%2Cw_1920%2Fformat%2Cjpeg&x-oss-signature=ouOhgsvDV0B%2Bqj2Vs7QrufooYbX7rUUmKuuJ1BFf2J8%3D&x-oss-signature-version=OSS2",
//"thumbnail": "https://bj29.cn-beijing.data.alicloudccp.com/f9n9Bwcf%2F1662258%2F605df29f93398d72d9cf487eba5f34ae3fb61637%2F605df29f031042c8b1714539afa2aa7b4e28feb5?x-oss-access-key-id=LTAIsE5mAn2F493Q&x-oss-additional-headers=referer&x-oss-expires=1630648496&x-oss-process=image%2Fresize%2Cw_400%2Fformat%2Cjpeg&x-oss-signature=8NXxKseSvy%2BR%2BsL92VkVFkk%2BxvrQOeAgZQIqV0uKOSY%3D&x-oss-signature-version=OSS2",
//"category": "image",
//"encrypt_mode": "none",
//"image_media_metadata": {
//"width": 1440,
//"height": 720,
//"image_tags": [
//{
//"confidence": 0.8713088035583496,
//"name": "艺术品",
//"tag_level": 1
//},
//{
//"confidence": 0.9999983310699463,
//"name": "日常用品",
//"tag_level": 1
//},
//{
//"confidence": 0.7486809492111206,
//"name": "其他场景",
//"tag_level": 1
//},
//{
//"confidence": 0.8712645769119263,
//"name": "其他事物",
//"tag_level": 1
//},
//{
//"confidence": 0.9999983310699463,
//"parent_name": "日常用品",
//"name": "文本",
//"tag_level": 2
//},
//{
//"confidence": 0.8712645769119263,
//"parent_name": "其他事物",
//"name": "小册子",
//"tag_level": 2
//},
//{
//"confidence": 0.8713088035583496,
//"parent_name": "艺术品",
//"name": "海报",
//"tag_level": 2
//},
//{
//"confidence": 0.824924111366272,
//"parent_name": "日常用品",
//"name": "图书",
//"tag_level": 2
//},
//{
//"confidence": 0.802035927772522,
//"parent_name": "日常用品",
//"name": "文件",
//"tag_level": 2
//},
//{
//"confidence": 0.7867134809494019,
//"parent_name": "日常用品",
//"name": "书本",
//"tag_level": 2
//},
//{
//"confidence": 0.7486809492111206,
//"parent_name": "其他场景",
//"name": "手机截图",
//"tag_level": 2
//},
//{
//"confidence": 0.9583567380905151,
//"parent_name": "日常用品",
//"name": "信",
//"tag_level": 2
//},
//{
//"confidence": 0.7308701276779175,
//"parent_name": "日常用品",
//"name": "手写",
//"tag_level": 2
//}
//],
//"exif": "{\"FileSize\":{\"value\":\"826909\"},\"Format\":{\"value\":\"jpg\"},\"ImageHeight\":{\"value\":\"720\"},\"ImageWidth\":{\"value\":\"1440\"},\"ResolutionUnit\":{\"value\":\"1\"},\"XResolution\":{\"value\":\"1/1\"},\"YResolution\":{\"value\":\"1/1\"}}",
//"image_quality": {
//"overall_score": 0.6130968332290649
//},
//"cropping_suggestion": [
//{
//"aspect_ratio": "1:1",
//"score": 0.6446176767349243,
//"cropping_boundary": {
//"width": 720,
//"height": 720,
//"top": 0,
//"left": 714
//}
//},
//{
//"aspect_ratio": "1:2",
//"score": 0.7018072009086609,
//"cropping_boundary": {
//"width": 360,
//"height": 720,
//"top": 0,
//"left": 714
//}
//},
//{
//"aspect_ratio": "23:16",
//"score": 0.7604787349700928,
//"cropping_boundary": {
//"width": 1035,
//"height": 720,
//"top": 0,
//"left": 178
//}
//},
//{
//"aspect_ratio": "2:1",
//"score": 0.7854302525520325,
//"cropping_boundary": {
//"width": 1440,
//"height": 720,
//"top": 0,
//"left": 0
//}
//},
//{
//"aspect_ratio": "2:3",
//"score": 0.6909570693969727,
//"cropping_boundary": {
//"width": 479,
//"height": 720,
//"top": 0,
//"left": 624
//}
//},
//{
//"aspect_ratio": "3:2",
//"score": 0.7645761370658875,
//"cropping_boundary": {
//"width": 1080,
//"height": 720,
//"top": 0,
//"left": 133
//}
//},
//{
//"aspect_ratio": "7:4",
//"score": 0.7802999019622803,
//"cropping_boundary": {
//"width": 1260,
//"height": 720,
//"top": 0,
//"left": 133
//}
//}
//]
//},
//"punish_flag": 0
//},

//视频
//{
//"drive_id": "1662258",
//"domain_id": "bj29",
//"file_id": "60e8ec2442706eab850443d09038298644ffeab2",
//"name": "Black.Widow.2021.2160p.DSNP.WEB-DL.DDP5.1.Atmos.HDR.HEVC-CMRG.mkv",
//"type": "file",
//"content_type": "application/oct-stream",
//"created_at": "2021-07-10T00:39:00.715Z",
//"updated_at": "2021-07-10T00:43:39.837Z",
//"user_meta": "{\"play_cursor\":\"722.410\"}",
//"parent_file_id": "root",
//"crc64_hash": "8340217521699108111",
//"content_hash": "5FAD6F55C2555D9D62E1AC8A7A736AEBC5EFC9DC",
//"content_hash_name": "sha1",
//"download_url": "https://bj29.cn-beijing.data.alicloudccp.com/4BA1D86x%2F248683%2F60e80365a652b873fffa49a09c38691d8a5d2a2f%2F60e80365c0ec610b26e248e2b777ea66c3d0c720?di=bj29&dr=1662258&f=60e8ec2442706eab850443d09038298644ffeab2&response-content-disposition=attachment%3B%20filename%2A%3DUTF-8%27%27Black.Widow.2021.2160p.DSNP.WEB-DL.DDP5.1.Atmos.HDR.HEVC-CMRG.mkv&u=ff08e5f172214ee88eb7790594e206b5&x-oss-access-key-id=LTAIsE5mAn2F493Q&x-oss-additional-headers=referer&x-oss-expires=1630648496&x-oss-signature=1VNkOiY8GLReG69XMBhj4OGy2X3HQoor6WtTld%2BDV0I%3D&x-oss-signature-version=OSS2",
//"url": "https://bj29.cn-beijing.data.alicloudccp.com/4BA1D86x%2F248683%2F60e80365a652b873fffa49a09c38691d8a5d2a2f%2F60e80365c0ec610b26e248e2b777ea66c3d0c720?di=bj29&dr=1662258&f=60e8ec2442706eab850443d09038298644ffeab2&u=ff08e5f172214ee88eb7790594e206b5&x-oss-access-key-id=LTAIsE5mAn2F493Q&x-oss-additional-headers=referer&x-oss-expires=1630648496&x-oss-signature=9MoZcuxzbkY%2BIdLq94Rsa41hdToSk%2FXm0zSb9mKJURo%3D&x-oss-signature-version=OSS2",
//"thumbnail": "https://bj29.cn-beijing.data.alicloudccp.com/4BA1D86x%2F248683%2F60e80365a652b873fffa49a09c38691d8a5d2a2f%2F60e80365c0ec610b26e248e2b777ea66c3d0c720?x-oss-access-key-id=LTAIsE5mAn2F493Q&x-oss-additional-headers=referer&x-oss-expires=1630648496&x-oss-process=video%2Fsnapshot%2Ct_120000%2Cf_jpg%2Cw_480%2Car_auto%2Cm_fast&x-oss-signature=RMZC1DTV7xyACJYkW2n0rJXok6DZIG8nLfimfSOz46Y%3D&x-oss-signature-version=OSS2",
//"category": "video",
//"encrypt_mode": "none",
//"video_media_metadata": {
//"time": "2021-07-09T06:52:32.000Z",
//"width": 3840,
//"height": 2160,
//"video_media_video_stream": [
//{
//"clarity": "2160",
//"fps": "24000/1001",
//"code_name": "hevc"
//}
//],
//"video_media_audio_stream": [
//{
//"channels": 6,
//"channel_layout": "",
//"code_name": "eac3",
//"sample_rate": "48000"
//}
//],
//"duration": "8026.784000"
//},
//"video_preview_metadata": {
//"bitrate": "17131680",
//"duration": "8026.784000",
//"audio_format": "eac3",
//"video_format": "hevc",
//"frame_rate": "24000/1001",
//"height": 2160,
//"width": 3840
//},
//"punish_flag": 0
//},

//{
//"drive_id": "1662258",
//"domain_id": "bj29",
//"file_id": "60794ad941ee2d8d24f843b7a0ffd80279927dfc",
//"name": "markdown",
//"type": "folder",
//"created_at": "2021-04-16T08:29:13.179Z",
//"updated_at": "2021-04-16T08:29:13.179Z",
//"hidden": false,
//"starred": false,
//"status": "available",
//"parent_file_id": "root",
//"encrypt_mode": "none"
//},
