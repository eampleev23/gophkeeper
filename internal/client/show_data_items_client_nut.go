package client

import (
	"encoding/json"
	"fmt"
	"github.com/DrSmithFr/go-console/question"
	"github.com/eampleev23/gophkeeper/internal/models"
	"net/http"
	"strconv"
)

var dataItems []models.DataItem

func (clientApp *ClientApp) ShowDataItems(response *http.Response) error {
	responseData, _, err := clientApp.RequestToApi(nil, "api/user/get-data-items", http.MethodGet)
	if err != nil {
		fmt.Println("Error while getting data-items")
		return err
	}
	err = json.Unmarshal(responseData, &dataItems)
	if err != nil {
		fmt.Println("Error while unmarshalling data-items")
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
		case 3:
			fmt.Printf("[текст]")
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
	if inputIDInt == 0 {
		// возврат на предыдущее меню
		clientApp.ShowAuthMenu(response)
		return nil
	}
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
	case "3":
		clientApp.ShowText(response, indexes[inputIDInt])
		break
	default:
		fmt.Println("Ошибка клиента, попробуйте обновить версию")
		return err
	}
	return nil
}
