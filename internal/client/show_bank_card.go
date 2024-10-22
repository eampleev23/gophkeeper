package client

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
	"net/http"
	"strings"
)

var bankCardItem models.BankCard

func (clientApp *ClientApp) ShowBankCard(response *http.Response, inputID string) error {

	var bankCardRequestStr = `{"id":`
	bankCardRequestStr += inputID
	bankCardRequestStr += `}`

	var bankCardRequest = []byte(bankCardRequestStr)

	responseData, _, err := clientApp.RequestToApi(bankCardRequest, "api/user/get-bank-card", http.MethodPost)
	if err != nil {
		fmt.Println("Ошибка при получении ответа от сервера, попробуйте обновить клиент")
		return err
	}
	err = json.Unmarshal(responseData, &bankCardItem)
	if err != nil {
		fmt.Println("Ошибка парсинга ответа сервера, попробуйте обновить клиент")
		return err
	}

	unPackedCardNumber := unpackCardNumber(bankCardItem)
	unPackedValidThru := unpackValidThru(bankCardItem)
	unPackedOwnerName := unpackOwnerName(bankCardItem)
	unPackedCVC := unpackCVC(bankCardItem)
	unPackedCVC = strings.TrimSuffix(unPackedCVC, ",")
	fmt.Printf("Номер карты: %s, Срок годности: %s\n", unPackedCardNumber, unPackedValidThru)
	fmt.Printf("Имя владельца: %s, Код CVC: %s\n", unPackedOwnerName, unPackedCVC)

	clientApp.ShowDataItems(nil)
	return nil
}

func unpackCardNumber(bankCardItem models.BankCard) (unpackedCardNumber string) {
	encryptedCardNumberBytes := convertMineToBytes(bankCardItem.CardNumber)
	encryptedNonceCardNumberBytes := convertMineToBytes(bankCardItem.NonceCardNumber)
	key := []byte("TuUdlQmYyD1DTaiGVV31ipyWnbKa0jUD")
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	unpackedCardNumberBytes, err := aesgcm.Open(nil, encryptedNonceCardNumberBytes, encryptedCardNumberBytes, nil) // расшифровываем
	if err != nil {
		fmt.Println("Ошибка клиента, попробуйте обновить версию")
		return
	}
	unpackedCardNumber = string(unpackedCardNumberBytes)
	return unpackedCardNumber
}

func unpackValidThru(bankCardItem models.BankCard) (unpackedValidThru string) {
	encryptedValidThruBytes := convertMineToBytes(bankCardItem.ValidThru)
	encryptedNonceValidThruBytes := convertMineToBytes(bankCardItem.NonceValidThru)
	key := []byte("TuUdlQmYyD1DTaiGVV31ipyWnbKa0jUD")
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	unpackedValidThruBytes, err := aesgcm.Open(nil, encryptedNonceValidThruBytes, encryptedValidThruBytes, nil) // расшифровываем
	if err != nil {
		fmt.Println("Ошибка клиента, попробуйте обновить версию")
		return
	}
	unpackedValidThru = string(unpackedValidThruBytes)
	return unpackedValidThru
}

func unpackOwnerName(bankCardItem models.BankCard) (unpackedOwnerName string) {
	encryptedOwnerNameBytes := convertMineToBytes(bankCardItem.OwnerName)
	encryptedNonceOwnerNameBytes := convertMineToBytes(bankCardItem.NonceOwnerName)
	key := []byte("TuUdlQmYyD1DTaiGVV31ipyWnbKa0jUD")
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	unpackedOwnerNameBytes, err := aesgcm.Open(nil, encryptedNonceOwnerNameBytes, encryptedOwnerNameBytes, nil) // расшифровываем
	if err != nil {
		fmt.Println("Ошибка клиента, попробуйте обновить версию")
		return
	}
	unpackedOwnerName = string(unpackedOwnerNameBytes)
	return unpackedOwnerName
}

func unpackCVC(bankCardItem models.BankCard) (unpackedCVC string) {
	encryptedCVCBytes := convertMineToBytes(bankCardItem.CVC)
	encryptedNonceCVCBytes := convertMineToBytes(bankCardItem.NonceCVC)
	key := []byte("TuUdlQmYyD1DTaiGVV31ipyWnbKa0jUD")
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	unpackedCVCBytes, err := aesgcm.Open(nil, encryptedNonceCVCBytes, encryptedCVCBytes, nil) // расшифровываем
	if err != nil {
		fmt.Println("Ошибка клиента, попробуйте обновить версию")
		return
	}
	unpackedCVC = string(unpackedCVCBytes)
	return unpackedCVC
}
