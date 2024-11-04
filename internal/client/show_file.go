package client

import (
	"fmt"
	"github.com/DrSmithFr/go-console/question"
	"io"
	"net/http"
	"os"
)

func (clientApp *ClientApp) ShowFile(response *http.Response, inputID string) error {

	pathToSaveFile := clientApp.Qh.Ask(
		question.
			NewQuestion(
				"Введите путь для сохранения файла: \n"),
	)

	fmt.Println(pathToSaveFile)

	var fileRequestStr = `{"id":9}`
	var fileRequest = []byte(fileRequestStr)

	// Create the file
	out, err := os.Create(pathToSaveFile)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	_, res, err := clientApp.RequestToApi(fileRequest, "api/user/get-file/", http.MethodGet)
	if err != nil {
		fmt.Println("Ошибка при получении ответа от сервера, попробуйте обновить клиент")
		return err
	}
	defer res.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	return nil
}
