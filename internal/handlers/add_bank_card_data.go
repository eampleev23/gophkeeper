package handlers

import "net/http"

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
}
