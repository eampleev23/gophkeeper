package main

import (
	"fmt"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/question"
	"net/http"
	"net/http/cookiejar"
	"os"
)

func main() {

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Ошибка клиента, попробуйте обновить версию..")
		main()
	}
	client := &http.Client{
		Jar: jar,
	}

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
		register(client, cmd, qh)
		break
	case "a":
		// запрашиваем данные для авторизации
		login(client, cmd, qh, nil)
		break
	}
}
