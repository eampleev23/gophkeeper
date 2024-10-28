package main

import (
	"github.com/eampleev23/gophkeeper/internal/handlers"
	"github.com/eampleev23/gophkeeper/internal/mlg"
	"github.com/eampleev23/gophkeeper/internal/myauth"
	"github.com/eampleev23/gophkeeper/internal/server_config"
	"github.com/eampleev23/gophkeeper/internal/services"
	"github.com/eampleev23/gophkeeper/internal/store"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var c *server_config.Config
var mL *mlg.ZapLog
var au *myauth.Authorizer
var s store.Store
var serv *services.Services
var h *handlers.Handlers

func TestRegister(t *testing.T) {

	c, err := server_config.NewConfig()
	if err != nil {
		t.Log(err)
	}

	mL, err := mlg.NewZapLogger(c.LogLevel)
	if err != nil {
		t.Log(err)
	}

	au, err := myauth.Initialize(c, mL)
	if err != nil {
		t.Log(err)
	}

	s, err := store.NewStorage(c, mL)
	if err != nil {
		t.Log(err)
	}

	serv := services.NewServices(s, c, mL, *au)
	h, err := handlers.NewHandlers(s, c, mL, *au, *serv)
	if err != nil {
		t.Log(err)
	}

	testCases := []struct {
		method       string
		expectedCode int
		contentType  string
	}{
		{method: http.MethodGet, expectedCode: http.StatusBadRequest, contentType: ""},
		{method: http.MethodPut, expectedCode: http.StatusBadRequest, contentType: ""},
		{method: http.MethodDelete, expectedCode: http.StatusBadRequest, contentType: ""},
		{method: http.MethodPost, expectedCode: http.StatusBadRequest, contentType: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, "/", nil)
			w := httptest.NewRecorder()

			// Вызовем хэндлер как обычную функцию без запуска сервера
			h.Register(w, r)
			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
		})
	}
}

func TestAuthentication(t *testing.T) {

	testCases := []struct {
		method       string
		expectedCode int
		contentType  string
	}{
		{method: http.MethodGet, expectedCode: http.StatusBadRequest, contentType: ""},
		{method: http.MethodPut, expectedCode: http.StatusBadRequest, contentType: ""},
		{method: http.MethodDelete, expectedCode: http.StatusBadRequest, contentType: ""},
		{method: http.MethodPost, expectedCode: http.StatusBadRequest, contentType: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, "/", nil)
			w := httptest.NewRecorder()

			// Вызовем хэндлер как обычную функцию без запуска сервера
			h.Authentication(w, r)
			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
		})
	}
}
