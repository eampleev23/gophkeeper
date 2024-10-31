package models

// TextDataItem - модель произвольных текстовых данных.
type TextDataItem struct {
	ID               int    `json:"id"`
	MetaValue        string `json:"meta-value"`
	TextContent      string `json:"text-content"`
	NonceTextContent string `json:"nonce_text_content"`
	OwnerID          int    `json:"owner-id"`
}

type TextDataItemReq struct {
	ID int `json:"id"`
}
