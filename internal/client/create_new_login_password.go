package client

import (
	"fmt"
	"github.com/DrSmithFr/go-console/question"
	"net/http"
)

func (clientApp *ClientApp) CreateNewLoginPassword(response *http.Response) error {

	metaValue := clientApp.Qh.Ask(
		question.
			NewQuestion(
				`Введите название (например "сайт facebook.com" или что-то в этом духе): `).
			SetDefaultAnswer("без названия"),
	)
	inputLogin := clientApp.Qh.Ask(
		question.
			NewQuestion(
				`Введите login: `),
	)
	inputPassword := clientApp.Qh.Ask(
		question.
			NewQuestion(
				`Введите Пароль: `),
	)
	var loginPassNewItemRequestStr = `{"meta-value": "`
	loginPassNewItemRequestStr += metaValue
	loginPassNewItemRequestStr += `", "login": "`
	loginPassNewItemRequestStr += inputLogin
	loginPassNewItemRequestStr += `", "password": "`
	loginPassNewItemRequestStr += inputPassword
	loginPassNewItemRequestStr += `,"}`

	var loginPassNewItemRequestBytes = []byte(loginPassNewItemRequestStr)

	_, response, err := clientApp.RequestToApi(loginPassNewItemRequestBytes, "api/user/add-login-password", http.MethodPost)
	if err != nil {
		fmt.Println("Ошибка при получении ответа от сервера, попробуйте обновить клиент")
		return err
	}
	isOk, err := clientApp.CheckStatusResponse(response)
	if isOk {
		fmt.Printf("Вы удачно добавили текстовые данные под названием %s \n", metaValue)
		clientApp.ShowAuthMenu(response)
		return nil
	} else {
		fmt.Printf("Ошибка клиента, попробуйте обновить приложение\n")
		clientApp.ShowAuthMenu(response)
		return nil
	}
	return nil
}
