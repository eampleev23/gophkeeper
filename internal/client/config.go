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
	RunAddr    string
	HttpClient *http.Client
	Cmd        *go_console.Script
	Qh         *question.Helper
	SecretKey  string
}

func NewApp() (*ClientApp, error) {

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Ошибка клиента, попробуйте обновить версию..")
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
	flag.StringVar(&clientApp.RunAddr, "a", "http://localhost:8080/", "Set listening address and port for server")

	// принимаем секретный ключ сервера для авторизации
	flag.StringVar(&clientApp.SecretKey, "s", "e4853f5c4810101e88f1898db21c15d3", "server's secret key for authorization")

	if envRunAddr := os.Getenv("RUN_ADDRESS"); envRunAddr != "" {
		clientApp.RunAddr = envRunAddr
	}

	if envSecretKey := os.Getenv("SECRET_KEY"); envSecretKey != "" {
		clientApp.SecretKey = envSecretKey
	}
	return nil
}
