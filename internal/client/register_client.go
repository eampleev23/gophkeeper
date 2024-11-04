package client

import (
	"fmt"
	"github.com/DrSmithFr/go-console/question"
	"net/http"
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
	_, response, err := clientApp.RequestToApi(registerRequest, "api/user/register", http.MethodPost)
	if err != nil {
		fmt.Println("Ошибка при получении ответа от сервера, попробуйте обновить клиент")
		return err
	}

	isOk, err := clientApp.CheckStatusResponse(response)
	if isOk {
		fmt.Printf("Вы удачно зарегистрировались и авторизовались %s \n")
		clientApp.ShowAuthMenu(response)
		return nil
	} else {
		fmt.Printf("Ошибка клиента, попробуйте обновить приложение\n")
		clientApp.ShowAuthMenu(response)
		return nil
	}
	return err
}
