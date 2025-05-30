package client

import (
	"flag"
	"fmt"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/question"
	"net/http"
	"net/http/cookiejar"
	"os"
)

type ClientApp struct {
	RunAddr          string
	HttpClient       *http.Client
	Cmd              *go_console.Script
	Qh               *question.Helper
	SecretKey        string
	SecretKeyForData string
}

func NewApp() (*ClientApp, error) {

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Ошибка клиента, попробуйте обновить версию..")
		return nil, err
	}

	client := &http.Client{
		Jar: jar,
	}

	cmd := go_console.NewScript().Build()
	qh := question.NewHelper(os.Stdin, cmd.Output)

	clientApp := &ClientApp{
		HttpClient: client,
		Cmd:        cmd,
		Qh:         qh,
	}

	err = clientApp.SetValues()
	if err != nil {
		return nil, err
	}
	return clientApp, nil
}

func (clientApp *ClientApp) SetValues() error {
	// регистрируем переменную flagRunAddr как аргумент -a со значением по умолчанию localhost:8080
	//flag.StringVar(&clientApp.RunAddr, "a", "http://localhost:8080/", "Set listening address and port for server")
	flag.StringVar(&clientApp.RunAddr, "a", "http://212.193.48.196:8080/", "Set listening address and port for server")

	// принимаем секретный ключ сервера для авторизации
	flag.StringVar(&clientApp.SecretKey, "s", "e4853f5c4810101e88f1898db21c15d3", "server's secret key for authorization")

	// принимаем секретный ключ сервера для расшифровки
	flag.StringVar(&clientApp.SecretKeyForData, "sd", "TuUdlQmYyD1DTaiGVV31ipyWnbKa0jUD", "secret key for data encrypting")

	if envRunAddr := os.Getenv("RUN_ADDRESS"); envRunAddr != "" {
		clientApp.RunAddr = envRunAddr
	}

	if envSecretKey := os.Getenv("SECRET_KEY"); envSecretKey != "" {
		clientApp.SecretKey = envSecretKey
	}

	if envSecretKeyForData := os.Getenv("SECRET_KEY_FOR_DATA"); envSecretKeyForData != "" {
		clientApp.SecretKeyForData = envSecretKeyForData
	}
	return nil
}
