package main

import (
	"encoding/json"
	"fmt"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/question"
	"github.com/eampleev23/gophkeeper/internal/models"
	"io"
	"net/http"
)

var dataItems []models.DataItem

func showDataItems(client *http.Client, cmd *go_console.Script, qh *question.Helper, response *http.Response) {
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/api/user/get-data-items", nil)
	if err != nil {
		fmt.Println("Ошибка формирования запроса, обратитесь к администратору")
	}
	response, err = client.Do(request)
	if err != nil {
		fmt.Println("Ошибка получения ответа, обратитесь к администратору")
	}
	if response.StatusCode == http.StatusOK {
		responseData, err := io.ReadAll(response.Body)
		err = json.Unmarshal(responseData, &dataItems)
		if err != nil {
			fmt.Println("ошибка:", err)
		}

		for _, dataItem := range dataItems {
			fmt.Printf("ID: %d - ", dataItem.ID)
			switch dataItem.DataType {
			case 1:
				fmt.Printf("[пара логин-пароль]")
				break
			default:
				fmt.Printf("[не известное значение, обновите клиент]")
			}

			fmt.Printf(" %s ", dataItem.MetaValue)
			fmt.Printf("\n")
		}
		inputID := qh.Ask(
			question.
				NewQuestion(
					"Для просмотра данных введите ID или 0 для возврата к предыдущему меню\n").
				SetDefaultAnswer("0"),
		)
		showDataItem(client, cmd, qh, response, inputID)
	}
}
