package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/question"
	"github.com/eampleev23/gophkeeper/internal/models"
	"io"
	"net/http"
	"strings"
)

var bankCardItem models.BankCard

func showBankCardItem(client *http.Client, cmd *go_console.Script, qh *question.Helper, response *http.Response, inputID string) {

	var bankCardRequestStr = `{"id":`
	bankCardRequestStr += inputID
	bankCardRequestStr += `}`

	var bankCardRequest = []byte(bankCardRequestStr)
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/user/get-bank-card", bytes.NewBuffer(bankCardRequest))
	if err != nil {
		fmt.Println("Ошибка формирования запроса, попробуйте обновить приложение")
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err = client.Do(request)
	if err != nil {
		fmt.Println("Ошибка формирования запроса, попробуйте обновить приложение")
	}
	if response.StatusCode == http.StatusOK {
		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Ошибка чтения ответа. попробуйте обновить клиент")
			return
		}
		err = json.Unmarshal(responseData, &bankCardItem)
		unPackedCardNumber := unpackCardNumber(bankCardItem)
		unPackedValidThru := unpackValidThru(bankCardItem)
		unPackedOwnerName := unpackOwnerName(bankCardItem)
		unPackedCVC := unpackCVC(bankCardItem)
		unPackedCVC = strings.TrimSuffix(unPackedCVC, ",")
		fmt.Printf("Номер карты: %s, Срок годности: %s\n", unPackedCardNumber, unPackedValidThru)
		fmt.Printf("Имя владельца: %s, Код CVC: %s\n", unPackedOwnerName, unPackedCVC)
	}
	showDataItems(client, cmd, qh, nil)
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
