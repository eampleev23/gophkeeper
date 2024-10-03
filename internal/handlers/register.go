package handlers

import (
	"net/http"
	"strings"
)

// Register регистрирует нового пользователя.
func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	// Проверяем формат запроса
	contentType := r.Header.Get("Content-Type")
	supportsJSON := strings.Contains(contentType, "application/json")
	if !supportsJSON {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
