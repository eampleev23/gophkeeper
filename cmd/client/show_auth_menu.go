package main

import (
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/question"
	"net/http"
)

func showAuthMenu(client *http.Client, cmd *go_console.Script, qh *question.Helper, response *http.Response) {
	secondMenuItem := qh.Ask(
		question.
			NewQuestion(
				"Для сохранения новой пары логин-пароль, введите l\n" +
					"Для выхода введите e\n"),
	)
	switch secondMenuItem {
	case "l":
		// собираем данные для новой пары логин-пароль
		createNewLoginPasswordItem(client, cmd, qh, response)
		break
	case "e":
		// разавторизуем пользователя и отправляем на приветственный экран
		logout(client, cmd, qh, response)
		break
	}
}
