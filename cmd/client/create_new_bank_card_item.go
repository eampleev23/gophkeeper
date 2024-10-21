package main

import (
	"bytes"
	"fmt"
	go_console "github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/question"
	"net/http"
)

func createNewBankCardItem(client *http.Client, cmd *go_console.Script, qh *question.Helper, response *http.Response) {
	inputMetaValue := qh.Ask(
		question.
			NewQuestion(
				`Введите название (например "кредитка альфабанк" или что-то в этом духе): `).
			SetDefaultAnswer("без названия"),
	)
	inputCardNumber := qh.Ask(
		question.
			NewQuestion(
				`Введите номер карты (например, "5521862555118522": `),
	)
	inputValidThru := qh.Ask(
		question.
			NewQuestion(
				`Введите срок действия карты (например, "12/28"): `),
	)
	inputOwnerName := qh.Ask(
		question.
			NewQuestion(
				`Введите имя и фамилию владельца карты (например, "EVGENII AMPLEEV"): `),
	)
	inputCVC := qh.Ask(
		question.
			NewQuestion(
				`Введите CVC-код карты с обратной стороны (например, "766"): `),
	)
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

	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/user/add-bank-card", bytes.NewBuffer(bankCardNewItemRequestBytes))
	if err != nil {
		fmt.Println("Ошибка, попробуйте обновить версию клиента")
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err = client.Do(request)
	if err != nil {
		fmt.Println("Ошибка получения ответа от сервера, попробуйте обновить версию клиента")
	}
	if response.StatusCode == http.StatusOK {
		fmt.Printf("Вы удачно добавили данные банковской карты под названием %s \n", inputMetaValue)
		showAuthMenu(client, cmd, qh, response)
	}
	if response.StatusCode != http.StatusBadRequest {
		fmt.Println("Ошибка получения ответа от сервера, попробуйте обновить версию клиента")
	}
	if response.StatusCode != http.StatusInternalServerError {
		fmt.Println("Ошибка получения ответа от сервера, попробуйте обновить версию клиента")
	}
	if response.StatusCode != http.StatusUnauthorized {
		fmt.Println("Необходимо авторизоваться")
		login(client, cmd, qh, response)
	}
}
