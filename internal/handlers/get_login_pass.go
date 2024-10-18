package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

/*
На вход хэндлер ожидает json такого формата:
{
    "id": "<id>",
}
*/

// GetLoginPass возвращает пользователю конкретные  логин-пароль.
func (h *Handlers) GetLoginPass(w http.ResponseWriter, r *http.Request) {
	h.l.ZL.Info("Handling GET Data Item")
	// Проверяем формат запроса
	contentType := r.Header.Get("Content-Type")
	supportsJSON := strings.Contains(contentType, "application/json")
	if !supportsJSON {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Проверяем авторизацию
	// Ппроверяем, не авторизован ли пользователь, отправивший запрос.
	h.l.ZL.Debug("Checking auth..")
	userID, isAuth, err := h.GetUserID(r)
	if err != nil {
		h.l.ZL.Error("GetUserID fail")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isAuth {
		h.l.ZL.Debug("Unauthorized user..")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	h.l.ZL.Debug("Authorized user:", zap.Int("userID", userID))

	// Получаем данные в случае корректного запроса.
	var req models.LoginPassReq
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		h.l.ZL.Info("cannot decode request JSON body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	loginPassItem, err := h.serv.GetLoginPassItemByID(r.Context(), userID, req.ID)
	if err != nil {
		h.l.ZL.Error("GetLoginPassItemByID fail", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println("loginPassItem.NoncePassword=", loginPassItem.NoncePassword)
	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")
	if err := enc.Encode(loginPassItem); err != nil {
		h.l.ZL.Info("fail encoding response in handler", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
