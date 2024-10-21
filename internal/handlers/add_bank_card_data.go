package handlers

import (
	"net/http"
	"strings"
)

/*
На вход хэндлер ожидает json такого формата(пользователь должен быть авторизован):
{
    "meta-value": "альфабанк кредитка",
    "card-number": "5521862555118522",
    "valid-thru": "10/29",
    "owner-name": "EVGENII AMPLEEV",
    "cvc": "987"
}
*/

func (h *Handlers) AddBankCardData(w http.ResponseWriter, r *http.Request) {
	h.l.ZL.Info("Handling /AddBankCardData request")
	// Проверяем формат запроса
	contentType := r.Header.Get("Content-Type")
	supportsJSON := strings.Contains(contentType, "application/json")
	if !supportsJSON {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Проверяем, не авторизован ли пользователь, отправивший запрос.
	ownerID, isAuth, err := h.GetUserID(r)
	if err != nil {
		h.l.ZL.Error("GetUserID fail")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isAuth {
		h.l.ZL.Info("Пользователь не авторизован, в ответе отказано")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	h.l.ZL.Info("Пользователь авторизован, можем двигаться дальше")

}
