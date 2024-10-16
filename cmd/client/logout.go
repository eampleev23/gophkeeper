package main

import (
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/question"
	"net/http"
)

func logout(client *http.Client, cmd *go_console.Script, qh *question.Helper, response *http.Response) {
	if response != nil {
		for _, v := range response.Cookies() {
			if v.Name == "token" {
				v.Name = "old"
				v.Value = "old"
				main()
				continue
			}
		}
		return
	}
}
