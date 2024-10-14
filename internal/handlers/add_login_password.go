package handlers

import "net/http"

/*
На вход хэндлер ожидает json такого формата(пользователь должен быть авторизован):
{
	"meta":{
		"name":"<name>"
		"desc":"<description>"
			},
    "login": "<login>",
    "password": "<password>"
}
*/

func (h *Handlers) AddLoginPasswordHandler(w http.ResponseWriter, r *http.Request) {
	h.l.ZL.Info("Add Login Password Handler")
}
