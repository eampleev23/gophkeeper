package handlers

import (
	"encoding/json"
	"github.com/eampleev23/gophkeeper/internal/models"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func (h *Handlers) GetFile(w http.ResponseWriter, r *http.Request) {
	h.l.ZL.Info("Handling GET file has called..")
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
	var req models.FileDataItemReq
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		h.l.ZL.Info("cannot decode request JSON body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fileItem, err := h.serv.GetFileItemByID(r.Context(), userID, req.ID)
	if err != nil {
		h.l.ZL.Error("GetLoginPassItemByID fail", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
	}
	h.l.ZL.Debug("got fileItem",
		zap.Int("id:", fileItem.ID),
		zap.String("server-path", fileItem.ServerPath),
	)
	// здесь нужно открыть этот файл и передать его на клиент, клиент создает у себя новый файл и записывает
	// в него данные из этого файла..
	http.ServeFile(w, r, fileItem.ServerPath)
}
