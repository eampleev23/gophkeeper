package main

// iter5 init
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/question"
	"github.com/eampleev23/gophkeeper/internal/models"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var logiPasswordItem models.LoginPassword

func showDataItem(client *http.Client, cmd *go_console.Script, qh *question.Helper, response *http.Response, inputID string) {
	var loginPassRequestStr = `{"id":`
	loginPassRequestStr += inputID
	loginPassRequestStr += `}`

	var loginPassRequest = []byte(loginPassRequestStr)
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/user/get-login-pass", bytes.NewBuffer(loginPassRequest))
	if err != nil {
		fmt.Println("Ошибка формирования запроса, обратитесь к администратору")
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err = client.Do(request)
	if err != nil {
		fmt.Println("Ошибка получения ответа, обратитесь к администратору")
	}
	if response.StatusCode == http.StatusOK {
		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Ошибка чтения ответа. попробуйте обновить клиент")
			return
		}
		err = json.Unmarshal(responseData, &logiPasswordItem)
	}

	unPackedLogin := unpackLogin(logiPasswordItem)
	unPackedPassword := unpackPassword(logiPasswordItem)
	fmt.Printf("Запрашиваемые логин и пароль: %s::%s\n", unPackedLogin, unPackedPassword)
	showDataItems(client, cmd, qh, nil)
}

func unpackLogin(inputLoginPassModel models.LoginPassword) (unpackedLogin string) {
	// все пришло в нормальном виде, теперь надо конвертнуть в байты
	encryptedLoginBytes := convertMineToBytes(inputLoginPassModel.Login)

	// теперь нужно получить байты из nonceLogin
	encryptedNonceLoginBytes := convertMineToBytes(inputLoginPassModel.NonceLogin)

	// теперь вроде бы есть все необходимое для расшифровки
	// расшифровываем
	key := []byte("TuUdlQmYyD1DTaiGVV31ipyWnbKa0jUD")

	// NewCipher создает и возвращает новый cipher.Block.
	// Ключевым аргументом должен быть ключ AES, 16, 24 или 32 байта
	// для выбора AES-128, AES-192 или AES-256.
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	// NewGCM возвращает заданный 128-битный блочный шифр
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	unpackedLoginBytes, err := aesgcm.Open(nil, encryptedNonceLoginBytes, encryptedLoginBytes, nil) // расшифровываем
	if err != nil {
		fmt.Println("Ошибка клиента, попробуйте обновить версию")
		return
	}
	// логин расшифровали корректно, ура!
	unpackedLogin = string(unpackedLoginBytes)
	return unpackedLogin
}

func unpackPassword(inputLoginPassModel models.LoginPassword) (unpackedPassword string) {
	// все пришло в нормальном виде, теперь надо конвертнуть в байты
	encryptedPasswordBytes := convertMineToBytes(inputLoginPassModel.Password)

	// теперь нужно получить байты из noncePassword
	encryptedNoncePasswordBytes := convertMineToBytes(inputLoginPassModel.NoncePassword)

	// теперь вроде бы есть все необходимое для расшифровки
	// расшифровываем
	key := []byte("TuUdlQmYyD1DTaiGVV31ipyWnbKa0jUD")

	// NewCipher создает и возвращает новый cipher.Block.
	// Ключевым аргументом должен быть ключ AES, 16, 24 или 32 байта
	// для выбора AES-128, AES-192 или AES-256.
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	// NewGCM возвращает заданный 128-битный блочный шифр
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	unpackedPasswordBytes, err := aesgcm.Open(nil, encryptedNoncePasswordBytes, encryptedPasswordBytes, nil) // расшифровываем
	if err != nil {
		fmt.Println("Ошибка клиента, попробуйте обновить версию")
		return
	}
	// логин расшифровали корректно, ура!
	unpackedPassword = string(unpackedPasswordBytes)
	unpackedPassword = strings.TrimSuffix(unpackedPassword, ",")
	return unpackedPassword
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
