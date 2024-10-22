package client_app

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
