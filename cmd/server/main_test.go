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

//
//func TestGetLoginPass(t *testing.T) {
//	handler := http.HandlerFunc(h.GetLoginPass)
//	srv := httptest.NewServer(handler)
//	defer srv.Close()
//
//	testCases := []struct {
//		name         string // добавляем название тестов
//		method       string
//		body         string // добавляем тело запроса в табличные тесты
//		expectedCode int
//		expectedBody string
//	}{
//		{
//			name:         "method_get",
//			method:       http.MethodGet,
//			expectedCode: http.StatusBadRequest,
//			expectedBody: "",
//		},
//		{
//			name:         "method_put",
//			method:       http.MethodPut,
//			expectedCode: http.StatusBadRequest,
//			expectedBody: "",
//		},
//		{
//			name:         "method_delete",
//			method:       http.MethodDelete,
//			expectedCode: http.StatusBadRequest,
//			expectedBody: "",
//		},
//	}
//	for _, tc := range testCases {
//		t.Run(tc.method, func(t *testing.T) {
//			req := resty.New().R()
//			req.Method = tc.method
//			req.URL = srv.URL
//
//			if len(tc.body) > 0 {
//				req.SetHeader("Content-Type", "application/json")
//				req.SetBody(tc.body)
//			}
//
//			resp, err := req.Send()
//			//assert.NoError(t, err, "error making HTTP request")
//			assert.NoError(t, err, "error making HTTP request")
//
//			assert.Equal(t, tc.expectedCode, resp.StatusCode(), "Response code didn't match expected")
//			// проверяем корректность полученного тела ответа, если мы его ожидаем
//			if tc.expectedBody != "" {
//				assert.JSONEq(t, tc.expectedBody, string(resp.Body()))
//			}
//		})
//	}
//}
