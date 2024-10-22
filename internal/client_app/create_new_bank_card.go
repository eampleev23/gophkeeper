package client_app

import (
	"bytes"
	"fmt"
	"github.com/DrSmithFr/go-console/question"
	"net/http"
	url2 "net/url"
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

	url, err := url2.JoinPath(clientApp.RunAddr, "api/user/add-bank-card")
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bankCardNewItemRequestBytes))
	if err != nil {
		fmt.Println("Ошибка, попробуйте обновить версию клиента")
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err = clientApp.HttpClient.Do(request)
	if err != nil {
		fmt.Println("Ошибка получения ответа от сервера, попробуйте обновить версию клиента")
	}
	if response.StatusCode == http.StatusOK {
		fmt.Printf("Вы удачно добавили данные банковской карты под названием %s \n", inputMetaValue)
		clientApp.ShowAuthMenu(response)
	}
	if response.StatusCode != http.StatusBadRequest {
		fmt.Println("Ошибка получения ответа от сервера, попробуйте обновить версию клиента")
	}
	if response.StatusCode != http.StatusInternalServerError {
		fmt.Println("Ошибка получения ответа от сервера, попробуйте обновить версию клиента")
	}
	if response.StatusCode != http.StatusUnauthorized {
		fmt.Println("Необходимо авторизоваться")
		clientApp.Login(response)
	}
	return nil
}
