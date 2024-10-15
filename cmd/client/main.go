package main

import (
	"fmt"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/question"
	"os"
)

func main() {
	cmd := go_console.NewScript().Build()
	qh := question.NewHelper(os.Stdin, cmd.Output)

	firstMenuItem := qh.Ask(
		question.
			NewQuestion(
				"Добро пожаловать в gophkeeper - менеджер паролей.\n" +
					"Для регистрации нового пользователя введите r\n" +
					"Для авторизации введите a\n").
			SetDefaultAnswer("r"),
	)
	switch firstMenuItem {
	case "r":
		// запрашиваем данные для регистрации
		register()
		break
	case "a":
		// запрашиваем данные для авторизации
		login()
		break
	}
}

func register() {
	cmd := go_console.NewScript().Build()
	qh := question.NewHelper(os.Stdin, cmd.Output)
	login := qh.Ask(
		question.
			NewQuestion(
				"Вы перешли к регистрации нового пользователя\n" +
					"Введите login:\n"),
	)
	password := qh.Ask(
		question.
			NewQuestion(
				"Какой пароль будете использовать при авторизации (при вводе он не будет показываться)?\n").SetHidden(true),
	)
	fmt.Println("Пробуем зарегистрировать вас с логином ", login, "и введенным паролем", password)
}

func login() {

}
