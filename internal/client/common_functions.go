package client

import (
	"fmt"
	"net/http"
)

func (clientApp *ClientApp) CheckStatusResponse(response *http.Response) (bool, error) {
	if response.StatusCode == http.StatusOK {
		return true, nil
	}
	if response.StatusCode == http.StatusBadRequest {
		fmt.Println("Ошибка получения ответа от сервера, попробуйте обновить версию клиента")
		return false, nil
	}
	if response.StatusCode == http.StatusInternalServerError {
		fmt.Println("Ошибка получения ответа от сервера, попробуйте обновить версию клиента")
		return false, nil
	}
	if response.StatusCode == http.StatusUnauthorized {
		fmt.Println("Необходимо авторизоваться")
		clientApp.Login(response)
		return false, nil
	}
	return false, nil
}
