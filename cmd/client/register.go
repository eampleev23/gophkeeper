package main

import (
	"bytes"
	"fmt"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/question"
	"net/http"
)

func register(client *http.Client, cmd *go_console.Script, qh *question.Helper) {
	inputLogin := qh.Ask(
		question.
			NewQuestion(
				"Вы перешли к регистрации нового пользователя\n" +
					"Введите login:\n"),
	)
	inputPassword := qh.Ask(
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
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/user/register", bytes.NewBuffer(registerRequest))
	if err != nil {
		fmt.Println("Ошибка формирования запроса, обратитесь к администратору")
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Ошибка получения ответа, обратитесь к администратору")
	}
	if response.StatusCode == http.StatusOK {
		fmt.Println("Вы удачно зарегистрировались и авторизовались")
		login(client, cmd, qh, response)
	}
	if response.StatusCode == http.StatusConflict {
		fmt.Println("Пользователь с таким логином уже зарегистрирован.")
		register(client, cmd, qh)
	}
	if response.StatusCode == http.StatusBadRequest {
		fmt.Println("Ошибка клиента, попробуйте обновить версию")
		main()
	}
	if response.StatusCode == http.StatusInternalServerError {
		fmt.Println("Внутренняя ошибка сервера, попробуйте еще раз..")
		register(client, cmd, qh)
	}
}
