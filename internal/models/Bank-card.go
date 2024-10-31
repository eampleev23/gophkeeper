package models

// BankCard - модель банковской карты.
type BankCard struct {
	ID              int    `json:"id"`
	MetaValue       string `json:"meta-value"`
	CardNumber      string `json:"card-number"`
	ValidThru       string `json:"valid-thru"`
	OwnerID         int    `json:"owner-id"`
	OwnerName       string `json:"owner-name"`
	CVC             string `json:"cvc"`
	NonceCardNumber string `json:"nonce-card-number"`
	NonceValidThru  string `json:"nonce-valid-thru"`
	NonceOwnerName  string `json:"nonce-owner-name"`
	NonceCVC        string `json:"nonce-cvc"`
}

type BancCardReq struct {
	ID int `json:"id"`
}
