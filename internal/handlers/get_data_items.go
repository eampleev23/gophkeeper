package handlers

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

// GetDataItems возвращает пользователю все его сохраненные пароли.
func (h *Handlers) GetDataItems(w http.ResponseWriter, r *http.Request) {
	h.l.ZL.Info("GetDataItems called")
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
	dataItems, err := h.serv.GetDataItemsByUserID(r.Context(), userID)
	if err != nil {
		h.l.ZL.Error("GetDataItemsByUserID fail", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
	}
	if len(dataItems) == 0 {
		h.l.ZL.Debug("No data for response")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")
	if err := enc.Encode(dataItems); err != nil {
		h.l.ZL.Info("fail encoding response in handler", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
