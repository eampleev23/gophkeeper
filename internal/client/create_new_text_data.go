package client

import (
	"fmt"
	"github.com/DrSmithFr/go-console/question"
	"github.com/icza/gox/stringsx"
	"net/http"
)

func (clientApp *ClientApp) CreateNewTextData(response *http.Response) error {
	inputMetaValue := clientApp.Qh.Ask(
		question.
			NewQuestion(
				`Введите название (например "моё кодовое слово в сбербанке" или что-то в этом духе): `).
			SetDefaultAnswer("без названия"),
	)
	inputTextContent := clientApp.Qh.Ask(
		question.
			NewQuestion(
				`Введите произвольный текст (например, "Кирсанова1965": `),
	)
	inputMetaValue = stringsx.Clean(inputMetaValue)
	inputTextContent = stringsx.Clean(inputTextContent)

	var textDataNewItemRequestStr = `{"meta-value": "`
	textDataNewItemRequestStr += inputMetaValue
	textDataNewItemRequestStr += `", "text-content": "`
	textDataNewItemRequestStr += inputTextContent
	textDataNewItemRequestStr += `,"}`

	var textDataNewItemRequestBytes = []byte(textDataNewItemRequestStr)
	_, response, err := clientApp.RequestToApi(textDataNewItemRequestBytes, "api/user/add-text", http.MethodPost)
	if err != nil {
		fmt.Println("Ошибка при получении ответа от сервера, попробуйте обновить клиент")
		return err
	}

	isOk, err := clientApp.CheckStatusResponse(response)
	if isOk {
		fmt.Printf("Вы удачно добавили текстовые данные под названием %s \n", inputMetaValue)
		clientApp.ShowAuthMenu(response)
		return nil
	} else {
		fmt.Printf("Ошибка клиента, попробуйте обновить приложение\n")
		clientApp.ShowAuthMenu(response)
		return nil
	}
}
