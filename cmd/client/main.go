package main

import (
	"bytes"
	"fmt"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/question"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"net/http/cookiejar"
	"os"
)

func main() {

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Ошибка клиента, попробуйте обновить версию..")
		main()
	}
	client := &http.Client{
		Jar: jar,
	}

	cmd := go_console.NewScript().Build()
	qh := question.NewHelper(os.Stdin, cmd.Output)

	firstMenuItem := qh.Ask(
		question.
			NewQuestion(
				"Добро пожаловать в gophkeeper - менеджер паролей.\n" +
					"Для регистрации нового пользователя введите r\n" +
					"Для авторизации введите a\n").
			SetDefaultAnswer("r"),
	)
	switch firstMenuItem {
	case "r":
		// запрашиваем данные для регистрации
		register(client, cmd, qh)
		break
	case "a":
		// запрашиваем данные для авторизации
		login(client, cmd, qh, nil)
		break
	}
}

func register(client *http.Client, cmd *go_console.Script, qh *question.Helper) {
	inputLogin := qh.Ask(
		question.
			NewQuestion(
				"Вы перешли к регистрации нового пользователя\n" +
					"Введите login:\n"),
	)
	inputPassword := qh.Ask(
		question.
			NewQuestion(
				"Какой пароль будете использовать при авторизации (при вводе он не будет показываться)?\n").SetHidden(true),
	)
	fmt.Println("Пробуем зарегистрировать вас с логином ", inputLogin, "и введенным паролем")

	var registerRequestStr = `{"login": "`
	registerRequestStr += inputLogin
	registerRequestStr += `", "password": "`
	registerRequestStr += inputPassword
	registerRequestStr += `,"}`

	var registerRequest = []byte(registerRequestStr)
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/user/register", bytes.NewBuffer(registerRequest))
	if err != nil {
		fmt.Println("Ошибка формирования запроса, обратитесь к администратору")
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Ошибка получения ответа, обратитесь к администратору")
	}
	if response.StatusCode == http.StatusOK {
		fmt.Println("Вы удачно зарегистрировались и авторизовались")
		login(client, cmd, qh, response)
	}
	if response.StatusCode == http.StatusConflict {
		fmt.Println("Пользователь с таким логином уже зарегистрирован.")
		register(client, cmd, qh)
	}
	if response.StatusCode == http.StatusBadRequest {
		fmt.Println("Ошибка клиента, попробуйте обновить версию")
		main()
	}
	if response.StatusCode == http.StatusInternalServerError {
		fmt.Println("Внутренняя ошибка сервера, попробуйте еще раз..")
		register(client, cmd, qh)
	}
}
func login(client *http.Client, cmd *go_console.Script, qh *question.Helper, response *http.Response) {
	// авторизован ли пользователь и если да, то под каким именем, приветствуем его и показываем меню
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
			continue
		}
	}
}

// Claims описывает утверждения, хранящиеся в токене + добавляет кастомное UserID.
type Claims struct {
	jwt.RegisteredClaims
	UserID    int
	UserLogin string
}
