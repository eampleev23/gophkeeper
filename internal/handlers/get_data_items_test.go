package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers_GetDataItems(t *testing.T) {
	// сначала описываем структуру - скелет ответа, который для нас важен в тесте хэндлера.
	type want struct {
		method string
		code   int
	}

	// далее описываем тест кейсы. Каждый тест кейс содержит имя и экземпляр структуры want с конкретными значениями
	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test #1", // первый кейс имеет имя positive test #1
			want: want{
				method: http.MethodGet,
				code:   http.StatusUnauthorized, // код ответа хэндлера 400
			},
		},
	}

	// перебираем все тест кейсы ( в нашем случае он всего один)
	for _, test := range tests { // значение будет храниться в переменной test
		t.Run(test.name, func(t *testing.T) { // запускаем в горутине каждый кейс параллельно
			// создаем реквест, используем httptest вместо http (выбрасывает панику вместо классического поведения)
			request := httptest.NewRequest(test.want.method, "/api/user/add-login-password", nil) // заполняем тип запроса, урл и тело.
			// создаём новый Recorder
			w := httptest.NewRecorder()
			handlers := initHandlersForTests(t)
			// вызываем хэндлер как обычную функцию
			handlers.GetDataItems(w, request)

			// записываем результат из w
			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
		})
	}
}
