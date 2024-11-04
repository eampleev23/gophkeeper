package main

import (
	"fmt"
	"github.com/DrSmithFr/go-console/question"
	"github.com/eampleev23/gophkeeper/internal/client"
	"log"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	app, err := client.NewApp()
	if err != nil {
		return fmt.Errorf("failed to initialize a new config: %w", err)
	}

	firstMenuItem := app.Qh.Ask(
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
		app.Register()
		break
	case "a":
		// запрашиваем данные для авторизации
		app.Login(nil)
		break
	}
	return nil
}
