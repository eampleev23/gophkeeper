package handlers

import (
	"bytes"
	"github.com/eampleev23/gophkeeper/internal/mlg"
	"github.com/eampleev23/gophkeeper/internal/myauth"
	"github.com/eampleev23/gophkeeper/internal/services"
	"github.com/eampleev23/gophkeeper/internal/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers_AddLoginPasswordData(t *testing.T) {

	mLAL, err := mlg.NewZapLogger(confForTests.LogLevel)
	if err != nil {
		t.Log(err)
	}

	auAL, err := myauth.Initialize(confForTests, mLAL)
	if err != nil {
		t.Log(err)
	}

	sAL, err := store.NewStorage(confForTests, mLAL)
	if err != nil {
		t.Log(err)
	}

	servAL := services.NewDBServices(sAL, confForTests, mLAL, *auAL)
	hAL, err := NewHandlers(sAL, confForTests, mLAL, *auAL, servAL)
	if err != nil {
		t.Log(err)
	}

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
			// вызываем хэндлер как обычную функцию
			hAL.AddLoginPasswordData(w, request)

			// записываем результат из w
			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			require.NoError(t, err)
		})
	}
}
