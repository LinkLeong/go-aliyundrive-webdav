package model

type FilePath struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	// 	created_at: "2021-09-06T07:12:29.103Z"
	// domain_id: "bj29"
	// drive_id: "1662258"
	// encrypt_mode: "none"
	// file_id: "6135bf5dfa3c0db759ce4f65b763eb895c422395"
	// hidden: false
	// name: "未命名文件夹"
	// parent_file_id: "6135bf58e00ab306ed024e75917eb91e179505b7"
	// starred: false
	// status: "available"
	// type: "folder"
	// updated_at: "2021-09-06T07:12:29.103Z"
}

type ListFilePath struct {
	Items []FilePath `json:"items"`
}
