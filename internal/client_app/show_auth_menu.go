package client_app

import (
	"github.com/DrSmithFr/go-console/question"
	"net/http"
)

func (clientApp *ClientApp) ShowAuthMenu(response *http.Response) {
	secondMenuItem := clientApp.Qh.Ask(
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
		//createNewLoginPasswordItem(client, cmd, qh, response)
		clientApp.CreateNewLoginPassword(response)
		break
	case "c":
		// собираем данные для новой банковской карты
		//createNewBankCardItem(client, cmd, qh, response)
		clientApp.CreateNewBankCard(response)
		break
	case "q":
		// разавторизуем пользователя и отправляем на приветственный экран
		//logout(client, cmd, qh, response)
		clientApp.Logout(response)
		break
	case "s":
		// показываем все сохраненные данные пользователя
		//showDataItems(client, cmd, qh, response)
		clientApp.ShowDataItems(response)
		break
	}
}
