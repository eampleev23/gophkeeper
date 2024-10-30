package client

import (
	"encoding/json"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
	"net/http"
	"strconv"
	"strings"
)

var loginPasswordItem models.LoginPassword

func (clientApp *ClientApp) ShowLoginPass(response *http.Response, inputID string) error {

	var loginPassRequestStr = `{"id":`
	loginPassRequestStr += inputID
	loginPassRequestStr += `}`

	var loginPassRequest = []byte(loginPassRequestStr)

	responseData, _, err := clientApp.RequestToApi(loginPassRequest, "api/user/get-login-pass", http.MethodPost)
	if err != nil {
		fmt.Println("Ошибка при получении ответа от сервера, попробуйте обновить клиент")
		return err
	}

	err = json.Unmarshal(responseData, &loginPasswordItem)

	unPackedLogin := UnpackData(loginPasswordItem.Login, loginPasswordItem.NonceLogin)
	unPackedPassword := UnpackData(loginPasswordItem.Password, loginPasswordItem.NoncePassword)
	unPackedPassword = strings.TrimSuffix(unPackedPassword, ",")
	fmt.Printf("\nЗапрашиваемые логин и пароль: %s::%s\n\n", unPackedLogin, unPackedPassword)

	clientApp.ShowDataItems(nil)
	return err
}

func convertMineToBytes(mineStr string) []byte {
	splitedStr := strings.Split(mineStr, "!")
	b := make([]byte, len(splitedStr)-1)
	for i := 0; i < len(splitedStr)-1; i++ {
		intVal, err := strconv.Atoi(splitedStr[i])
		if err != nil {
			fmt.Println("Ошибка клиента, попробуйте обновить версию")
			return nil
		}
		b[i] = byte(intVal)
	}
	return b
}
