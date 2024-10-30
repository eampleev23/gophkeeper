package client

import (
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

	unPackedCardNumber := UnpackData(bankCardItem.CardNumber, bankCardItem.NonceCardNumber)
	unPackedValidThru := UnpackData(bankCardItem.ValidThru, bankCardItem.NonceValidThru)
	unPackedOwnerName := UnpackData(bankCardItem.OwnerName, bankCardItem.NonceOwnerName)
	unPackedCVC := UnpackData(bankCardItem.CVC, bankCardItem.NonceCVC)
	unPackedCVC = strings.TrimSuffix(unPackedCVC, ",")
	fmt.Printf("\nНомер карты: %s, Срок годности: %s,\n", unPackedCardNumber, unPackedValidThru)
	fmt.Printf("Имя владельца: %s, Код CVC: %s.\n\n", unPackedOwnerName, unPackedCVC)

	clientApp.ShowDataItems(nil)
	return nil
}
