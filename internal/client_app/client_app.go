package client_app

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

	clientApp.SetValues()
	return clientApp, nil
}

func (clientApp *ClientApp) SetValues() {
	// регистрируем переменную flagRunAddr как аргумент -a со значением по умолчанию localhost:8080
	flag.StringVar(&clientApp.RunAddr, "a", "http://localhost:8080/", "Set listening address and port for server")

	if envRunAddr := os.Getenv("RUN_ADDRESS"); envRunAddr != "" {
		clientApp.RunAddr = envRunAddr
	}
}
