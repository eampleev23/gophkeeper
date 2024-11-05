package client

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func (clientApp *ClientApp) UnpackData(packedStr, nonceStr string) (unpackedStr string) {
	encryptedStrBytes := convertMineToBytes(packedStr)
	encryptedNonceBytes := convertMineToBytes(nonceStr)
	key := []byte(clientApp.SecretKeyForData)
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	// NewGCM возвращает заданный 128-битный блочный шифр
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	unpackedStrBytes, err := aesgcm.Open(nil, encryptedNonceBytes, encryptedStrBytes, nil) // расшифровываем
	if err != nil {
		fmt.Println("Ошибка клиента, попробуйте обновить версию")
		return
	}
	unpackedStr = string(unpackedStrBytes)
	return unpackedStr
}
