package handlers

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers_AddLoginPasswordData(t *testing.T) {

	// сначала описываем структуру - скелет ответа, который для нас важен в тесте хэндлера.
	type want struct {
		code int
	}

	// далее описываем тест кейсы. Каждый тест кейс содержит имя и экземпляр структуры want с конкретными значениями
	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test #1", // первый кейс имеет имя positive test #1
			want: want{
				code: 400, // код ответа хэндлера 200
			},
		},
	}

	// перебираем все тест кейсы ( в нашем случае он всего один)
	for _, test := range tests { // значение будет храниться в переменной test
		t.Run(test.name, func(t *testing.T) { // запускаем в горутине каждый кейс параллельно
			// создаем реквест, используем httptest вместо http (выбрасывает панику вместо классического поведения)

			body := []byte(`{
  "meta-value":"сайт vk",
  "login": "+79261382009",
  "password": "пароль"
}`)
			var ioReader io.Reader
			ioReader = bytes.NewReader(body)
			request := httptest.NewRequest(http.MethodPost, "/api/user/add-login-password", ioReader) // заполняем тип запроса, урл и тело.
			// создаём новый Recorder
			w := httptest.NewRecorder()
			handlers := initHandlersForTests(t)
			// вызываем хэндлер как обычную функцию
			handlers.AddLoginPasswordData(w, request)

			// записываем результат из w
			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
		})
	}
}
