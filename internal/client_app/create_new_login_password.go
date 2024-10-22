package client_app

import (
	"bytes"
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

	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/user/add-login-password", bytes.NewBuffer(loginPassNewItemRequestBytes))
	if err != nil {
		fmt.Println("Ошибка, попробуйте обновить версию клиента")
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err = clientApp.HttpClient.Do(request)
	if err != nil {
		fmt.Println("Ошибка получения ответа от сервера, попробуйте обновить версию клиента")
	}
	if response.StatusCode == http.StatusOK {
		fmt.Printf("Вы удачно добавили пару логин-пароль под названием %s \n", metaValue)
		//showAuthMenu(client, cmd, qh, response)
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
		//login(client, cmd, qh, response)
		clientApp.Login(response)
	}
	return nil
}
