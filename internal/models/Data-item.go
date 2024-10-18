package models

// DataItem - модель сохраненного типа данных.
type DataItem struct {
	ID        int    `json:"id"`
	DataType  int    `json:"data_type""`
	OwnerID   int    `json:"owner_id"`
	MetaValue string `json:"meta_value"`
}
