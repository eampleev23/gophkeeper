package models

// LoginPassword - модель пары логин-пароль.
type LoginPassword struct {
	ID            int    `json:"id"`
	MetaName      string `json:"meta-name"`
	MetaValue     string `json:"meta-value"`
	Login         string `json:"login"`
	Password      string `json:"password"`
	OwnerID       int    `json:"owner-id"`
	NonceLogin    string `json:"nonce-login"`
	NoncePassword string `json:"nonce-password"`
}

type LoginPassReq struct {
	ID int `json:"id"`
}
