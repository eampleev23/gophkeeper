package client

import (
	"bytes"
	"fmt"
	"github.com/DrSmithFr/go-console/question"
	"github.com/icza/gox/stringsx"
	"io"
	"mime/multipart"
	"net/http"
	url2 "net/url"
	"os"
)

func (clientApp *ClientApp) CreateNewFile(response *http.Response) error {
	inputMetaValue := clientApp.Qh.Ask(
		question.
			NewQuestion(
				`Введите название для загружаемого файла:`).
			SetDefaultAnswer("без названия"),
	)
	inputFilePath := clientApp.Qh.Ask(
		question.
			NewQuestion(
				`Введите путь к файлу, который необходимо загрузить на сервер (не более 10 мб)`),
	)
	inputMetaValue = stringsx.Clean(inputMetaValue)
	inputFilePath = stringsx.Clean(inputFilePath)

	file, _ := os.Open(inputFilePath) // открываем файл
	defer file.Close()                // не забываем закрыть

	body := &bytes.Buffer{} // создаём буфер

	// на основе буфера конструируем multipart.Writer из пакета mime/multipart
	writer := multipart.NewWriter(body)

	// готовим форму для отправки файла на сервер
	part, err := writer.CreateFormFile("uploadfile", inputFilePath)
	if err != nil {
		return err
	}

	// копируем файл в форму
	// multipart.Writer отформатирует данные и запишет в предоставленный буфер
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	writer.Close()
	handler := "/api/user/add-file"
	url, err := url2.JoinPath(clientApp.RunAddr, handler)
	if err != nil {

		return err
	}
	// пишем запрос
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}
	// добавляем заголовок запроса
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Add("Meta-Value", inputMetaValue)
	response, err = clientApp.HttpClient.Do(request)
	if err != nil {
		return err
	}
	fmt.Printf("Вы удачно добавили файл под названием %s \n", inputMetaValue)
	err = clientApp.ShowDataItems(response)
	if err != nil {
		return err
	}
	return nil
}
