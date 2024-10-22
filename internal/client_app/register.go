package client_app

import (
	"bytes"
	"fmt"
	"github.com/DrSmithFr/go-console/question"
	"net/http"
	url2 "net/url"
)

func (clientApp *ClientApp) Register() error {
	inputLogin := clientApp.Qh.Ask(
		question.
			NewQuestion(
				"Вы перешли к регистрации нового пользователя\n" +
					"Введите login:\n"),
	)
	inputPassword := clientApp.Qh.Ask(
		question.
			NewQuestion(
				"Какой пароль будете использовать при авторизации (при вводе он не будет показываться)?\n").SetHidden(true),
	)
	fmt.Println("Пробуем зарегистрировать вас с логином ", inputLogin, "и введенным паролем")

	var registerRequestStr = `{"login": "`
	registerRequestStr += inputLogin
	registerRequestStr += `", "password": "`
	registerRequestStr += inputPassword
	registerRequestStr += `,"}`

	var registerRequest = []byte(registerRequestStr)
	url, err := url2.JoinPath(clientApp.RunAddr, "api/user/register")
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(registerRequest))

	if err != nil {
		fmt.Println("Ошибка формирования запроса, обратитесь к администратору")
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err := clientApp.HttpClient.Do(request)
	if err != nil {
		fmt.Println("Ошибка получения ответа, обратитесь к администратору")
		return err
	}
	if response.StatusCode == http.StatusOK {
		fmt.Println("Вы удачно зарегистрировались и авторизовались")
		clientApp.Login(response)
		return nil
	}
	if response.StatusCode == http.StatusConflict {
		fmt.Println("Пользователь с таким логином уже зарегистрирован.")
		clientApp.Register()
		return err
	}
	if response.StatusCode == http.StatusBadRequest {
		fmt.Println("Ошибка клиента, попробуйте обновить версию")
		clientApp.Register()
		return err
	}
	if response.StatusCode == http.StatusInternalServerError {
		fmt.Println("Внутренняя ошибка сервера, попробуйте еще раз..")
		clientApp.Register()
		return err
	}
	return err
}
