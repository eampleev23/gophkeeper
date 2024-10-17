package main

import (
	"bytes"
	"fmt"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/question"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func login(client *http.Client, cmd *go_console.Script, qh *question.Helper, response *http.Response) {
	// авторизован ли пользователь и если да, то под каким именем, приветствуем его и показываем меню
	if response != nil {
		for _, v := range response.Cookies() {
			if v.Name == "token" {
				// значит клиент уже авторизован, можем получить его логин и поприветствовать
				// Создаем экземпляр структуры с утверждениями
				claims := &Claims{}
				// Парсим из строки токена tokenString в структуру claims
				_, err := jwt.ParseWithClaims(v.Value, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte("e4853f5c4810101e88f1898db21c15d3"), nil
				})
				if err != nil {
					fmt.Println("Ошибка авторизации, обратитесь к администратору")
				}
				fmt.Printf("Добро пожаловать в gophkeeper, %s!\n", claims.UserLogin)
				showAuthMenu(client, cmd, qh, response)
				continue
			}
		}
		return
	}
	inputLogin := qh.Ask(
		question.
			NewQuestion("Введите login\n"),
	)
	inputPassword := qh.Ask(
		question.
			NewQuestion("Введите пароль (он не будет отображаться в процессе ввода)\n").SetHidden(true),
	)
	var loginRequestStr = `{"login": "`
	loginRequestStr += inputLogin
	loginRequestStr += `", "password": "`
	loginRequestStr += inputPassword
	loginRequestStr += `,"}`

	var loginRequest = []byte(loginRequestStr)
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/user/login", bytes.NewBuffer(loginRequest))
	if err != nil {
		fmt.Println("Ошибка, попробуйте обновить версию клиента")
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err = client.Do(request)
	if err != nil {
		fmt.Println("Ошибка получения ответа, попробуйте обновить версию клиента")
	}

	if response.StatusCode == http.StatusOK {
		fmt.Println("Вы удачно авторизовались")
		login(client, cmd, qh, response)
	}
	if response.StatusCode == http.StatusUnauthorized {
		fmt.Println("Пользователь с такими логином и паролем не зарегистрирован")
		main()
	}
	if response.StatusCode == http.StatusBadRequest {
		fmt.Println("Ошибка клиента, попробуйте обновить версию")
		main()
	}
	if response.StatusCode == http.StatusInternalServerError {
		fmt.Println("Внутренняя ошибка сервера, попробуйте еще раз..")
		login(client, cmd, qh, nil)
	}
}
