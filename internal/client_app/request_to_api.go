package client_app

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	url2 "net/url"
)

func (clientApp *ClientApp) RequestToApi(jsonData []byte, handler string, method string) (responseData []byte, response *http.Response, err error) {

	url, err := url2.JoinPath(clientApp.RunAddr, handler)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка формирования запроса, попробуйте обновить клиент")
		return nil, nil, err
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err = clientApp.HttpClient.Do(request)
	if err != nil {
		fmt.Println("Ошибка получения ответа, обратитесь к администратору")
		return nil, nil, err
	}

	if response.StatusCode == http.StatusOK {
		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Ошибка чтения ответа. Попробуйте обновить клиент")
			return nil, nil, err
		}
		return responseData, response, nil
	}
	return nil, response, errors.New(string(response.StatusCode))
}
