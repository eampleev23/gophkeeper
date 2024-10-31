package models

// FileDataItem - модель файла.
type FileDataItem struct {
	ID         int    `json:"id"`
	MetaValue  string `json:"meta-value"`
	ClientPath string `json:"client-path"`
	ServerPath string `json:"server-path"`
	OwnerID    int    `json:"owner-id"`
}

type FileDataItemReq struct {
	ID int `json:"id"`
}
