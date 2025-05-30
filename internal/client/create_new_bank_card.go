package client

import (
	"fmt"
	"github.com/DrSmithFr/go-console/question"
	"github.com/icza/gox/stringsx"
	"net/http"
)

func (clientApp *ClientApp) CreateNewBankCard(response *http.Response) error {
	inputMetaValue := clientApp.Qh.Ask(
		question.
			NewQuestion(
				`Введите название (например "кредитка альфабанк" или что-то в этом духе): `).
			SetDefaultAnswer("без названия"),
	)
	inputCardNumber := clientApp.Qh.Ask(
		question.
			NewQuestion(
				`Введите номер карты (например, "5521862555118522": `),
	)
	inputValidThru := clientApp.Qh.Ask(
		question.
			NewQuestion(
				`Введите срок действия карты (например, "12/28"): `),
	)
	inputOwnerName := clientApp.Qh.Ask(
		question.
			NewQuestion(
				`Введите имя и фамилию владельца карты (например, "EVGENII AMPLEEV"): `),
	)
	inputCVC := clientApp.Qh.Ask(
		question.
			NewQuestion(
				`Введите CVC-код карты с обратной стороны (например, "766"): `),
	)

	inputCardNumber = stringsx.Clean(inputCardNumber)
	inputValidThru = stringsx.Clean(inputValidThru)
	inputOwnerName = stringsx.Clean(inputOwnerName)
	inputCVC = stringsx.Clean(inputCVC)

	var bankCardNewItemRequestStr = `{"meta-value": "`
	bankCardNewItemRequestStr += inputMetaValue
	bankCardNewItemRequestStr += `", "card-number": "`
	bankCardNewItemRequestStr += inputCardNumber
	bankCardNewItemRequestStr += `", "valid-thru": "`
	bankCardNewItemRequestStr += inputValidThru
	bankCardNewItemRequestStr += `", "owner-name": "`
	bankCardNewItemRequestStr += inputOwnerName
	bankCardNewItemRequestStr += `", "cvc": "`
	bankCardNewItemRequestStr += inputCVC
	bankCardNewItemRequestStr += `,"}`

	var bankCardNewItemRequestBytes = []byte(bankCardNewItemRequestStr)

	_, response, err := clientApp.RequestToApi(bankCardNewItemRequestBytes, "api/user/add-bank-card", http.MethodPost)
	if err != nil {
		fmt.Println("Ошибка при получении ответа от сервера, попробуйте обновить клиент")
		return err
	}
	isOk, err := clientApp.CheckStatusResponse(response)
	if isOk {
		fmt.Printf("Вы удачно добавили данные банковской карты %s \n", inputMetaValue)
		clientApp.ShowAuthMenu(response)
		return nil
	} else {
		fmt.Printf("Ошибка клиента, попробуйте обновить приложение\n")
		clientApp.ShowAuthMenu(response)
		return nil
	}
}
