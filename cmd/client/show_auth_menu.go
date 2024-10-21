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
					"Для сохранения новых данных банковской карты, введите c\n" +
					"Для просмотра сохраненных данных, введите s\n" +
					"Для выхода введите q\n"),
	)
	switch secondMenuItem {
	case "l":
		// собираем данные для новой пары логин-пароль
		createNewLoginPasswordItem(client, cmd, qh, response)
		break
	case "c":
		// собираем данные для новой банковской карты
		createNewBankCardItem(client, cmd, qh, response)
		break
	case "q":
		// разавторизуем пользователя и отправляем на приветственный экран
		logout(client, cmd, qh, response)
		break
	case "s":
		// показываем все сохраненные данные пользователя
		showDataItems(client, cmd, qh, response)
		break
	}

}
