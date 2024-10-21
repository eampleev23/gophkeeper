package handlers

import (
	"encoding/json"
	"github.com/eampleev23/gophkeeper/internal/models"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

/*
На вход хэндлер ожидает json такого формата(пользователь должен быть авторизован):
{
    "meta-name":"сайт",
    "meta-value": "lostfilm.tv",
    "login": "eampleev@yandex.ru",
    "password": "Newpass34,"
}
*/

func (h *Handlers) AddLoginPasswordData(w http.ResponseWriter, r *http.Request) {
	h.l.ZL.Info("Add Login Password Handler")
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
	// Получаем данные в случае корректного запроса.
	var inputModel models.LoginPassword
	// Декодер работает потоково, кажется это правильнее + короче, чем анмаршал.
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&inputModel); err != nil {
		h.l.ZL.Info("cannot decode request JSON body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// записываем значение ид автора запроса
	inputModel.OwnerID = ownerID
	outputModel, err := h.serv.InsertLoginPassword(r.Context(), inputModel)
	if err != nil {
		h.l.ZL.Error("h.serv.InsertLoginPassword fail..", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Если мы здесь, то логин-пароль успешно добавлены.
	h.l.ZL.Info("Success creating new login password", zap.Any("outputModel", outputModel))
	w.WriteHeader(http.StatusOK)
	return
}
