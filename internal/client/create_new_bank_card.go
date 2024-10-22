package client

import (
	"fmt"
	"github.com/DrSmithFr/go-console/question"
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
	if response.StatusCode == http.StatusOK {
		fmt.Printf("Вы удачно добавили данные банковской карты под названием %s \n", inputMetaValue)
		clientApp.ShowAuthMenu(response)
		return nil
	}
	if response.StatusCode != http.StatusBadRequest {
		fmt.Println("Ошибка получения ответа от сервера, попробуйте обновить версию клиента")
		return err
	}
	if response.StatusCode != http.StatusInternalServerError {
		fmt.Println("Ошибка получения ответа от сервера, попробуйте обновить версию клиента")
		return err
	}
	if response.StatusCode != http.StatusUnauthorized {
		fmt.Println("Необходимо авторизоваться")
		clientApp.Login(response)
		return err
	}
	return err
}
