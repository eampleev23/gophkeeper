package client_app

import (
	"bytes"
	"fmt"
	"github.com/DrSmithFr/go-console/question"
	"net/http"
)

func (clientApp *ClientApp) Register() {
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
	//url:= url2.JoinPath("")
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/user/register", bytes.NewBuffer(registerRequest))

	if err != nil {
		fmt.Println("Ошибка формирования запроса, обратитесь к администратору")
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err := clientApp.HttpClient.Do(request)
	if err != nil {
		fmt.Println("Ошибка получения ответа, обратитесь к администратору")
	}
	if response.StatusCode == http.StatusOK {
		fmt.Println("Вы удачно зарегистрировались и авторизовались")
		clientApp.Login(response)
	}
	if response.StatusCode == http.StatusConflict {
		fmt.Println("Пользователь с таким логином уже зарегистрирован.")
		clientApp.Register()
	}
	if response.StatusCode == http.StatusBadRequest {
		fmt.Println("Ошибка клиента, попробуйте обновить версию")
		clientApp.Register()
	}
	if response.StatusCode == http.StatusInternalServerError {
		fmt.Println("Внутренняя ошибка сервера, попробуйте еще раз..")
		clientApp.Register()
	}
}
