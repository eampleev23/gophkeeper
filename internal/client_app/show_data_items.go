package client_app

import (
	"encoding/json"
	"fmt"
	"github.com/DrSmithFr/go-console/question"
	"github.com/eampleev23/gophkeeper/internal/models"
	"io"
	"net/http"
	url2 "net/url"
	"strconv"
)

var dataItems []models.DataItem

func (clientApp *ClientApp) ShowDataItems(response *http.Response) error {
	url, err := url2.JoinPath(clientApp.RunAddr, "api/user/get-data-items")
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Ошибка формирования запроса, обратитесь к администратору")
		return err
	}
	response, err = clientApp.HttpClient.Do(request)
	if err != nil {
		fmt.Println("Ошибка получения ответа, обратитесь к администратору")
		return err
	}
	if response.StatusCode == http.StatusOK {
		responseData, err := io.ReadAll(response.Body)
		err = json.Unmarshal(responseData, &dataItems)
		if err != nil {
			fmt.Println("ошибка:", err)
			return err
		}

		var indexes = make(map[int]string)
		indexFront := 1
		for _, dataItem := range dataItems {
			fmt.Printf("ID: [%d] - ", indexFront)
			switch dataItem.DataType {
			case 1:
				fmt.Printf("[пара логин-пароль]")
				break
			case 2:
				fmt.Printf("[банковская карта]")
				break
			default:
				fmt.Printf("[не известное значение, обновите клиент]")
			}

			fmt.Printf(" -- %s -- ", dataItem.MetaValue)
			fmt.Printf(" дата добавления: %s", dataItem.CreatedAt.Format("02.01.2006 15:04:05"))
			fmt.Printf("\n")
			indexes[indexFront] = strconv.Itoa(dataItem.ID)
			indexFront++
		}
		inputID := clientApp.Qh.Ask(
			question.
				NewQuestion(
					"Для просмотра данных введите ID или 0 для возврата к предыдущему меню\n").
				SetDefaultAnswer("0"),
		)
		inputIDInt, err := strconv.Atoi(inputID)
		if err != nil {
			fmt.Println("Ошибка клиента, попробуйте обновить версию")
			return err
		}
		// определяем какого типа зашифрованное значение хочет посмотреть пользователь (логин пароль или банковская карта и тд)
		var typeOfDataToShow string
		for _, dataItem := range dataItems {
			dataItemIdInt := strconv.Itoa(dataItem.ID)
			if dataItemIdInt == indexes[inputIDInt] {
				typeOfDataToShow = strconv.Itoa(dataItem.DataType)
			}
		}
		switch typeOfDataToShow {
		case "1":
			clientApp.ShowLoginPass(response, indexes[inputIDInt])
			break
		case "2":
			clientApp.ShowBankCard(response, indexes[inputIDInt])
			break
		default:
			fmt.Println("Ошибка клиента, попробуйте обновить версию")
			return err
		}
	}
	return err
}
