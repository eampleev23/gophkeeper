package client_app

import (
	"bytes"
	"fmt"
	"github.com/DrSmithFr/go-console/question"
	"github.com/eampleev23/gophkeeper/internal/myauth"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	url2 "net/url"
)

func (clientApp *ClientApp) Login(response *http.Response) error {
	// авторизован ли пользователь и если да, то под каким именем, приветствуем его и показываем меню
	if response != nil {
		for _, v := range response.Cookies() {
			if v.Name == "token" {
				// значит клиент уже авторизован, можем получить его логин и поприветствовать
				// Создаем экземпляр структуры с утверждениями
				claims := &myauth.Claims{}
				// Парсим из строки токена tokenString в структуру claims
				_, err := jwt.ParseWithClaims(v.Value, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(clientApp.SecretKey), nil
				})
				if err != nil {
					fmt.Println("Ошибка авторизации, обратитесь к администратору")
				}
				fmt.Printf("Добро пожаловать в gophkeeper, %s!\n", claims.UserLogin)
				clientApp.ShowAuthMenu(response)
				continue
			}
		}
		return nil
	}
	inputLogin := clientApp.Qh.Ask(
		question.
			NewQuestion("Введите login\n"),
	)
	inputPassword := clientApp.Qh.Ask(
		question.
			NewQuestion("Введите пароль (он не будет отображаться в процессе ввода)\n").SetHidden(true),
	)
	var loginRequestStr = `{"login": "`
	loginRequestStr += inputLogin
	loginRequestStr += `", "password": "`
	loginRequestStr += inputPassword
	loginRequestStr += `,"}`

	var loginRequest = []byte(loginRequestStr)
	url, err := url2.JoinPath(clientApp.RunAddr, "api/user/login")
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(loginRequest))
	if err != nil {
		fmt.Println("Ошибка, попробуйте обновить версию клиента")
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err = clientApp.HttpClient.Do(request)
	if err != nil {
		fmt.Println("Ошибка получения ответа, попробуйте обновить версию клиента")
		return err
	}
	if response.StatusCode == http.StatusOK {
		fmt.Println("Вы удачно авторизовались")
		clientApp.Login(response)
		return nil
	}
	if response.StatusCode == http.StatusUnauthorized {
		fmt.Println("Пользователь с такими логином и паролем не зарегистрирован")
		return err
	}
	if response.StatusCode == http.StatusBadRequest {
		fmt.Println("Ошибка клиента, попробуйте обновить версию")
		return err
	}
	if response.StatusCode == http.StatusInternalServerError {
		fmt.Println("Внутренняя ошибка сервера, попробуйте еще раз..")
		clientApp.Login(nil)
		return err
	}
	return err
}
