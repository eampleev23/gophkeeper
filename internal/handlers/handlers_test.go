package handlers

import (
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

var handlersForTests *Handlers

func initHandlersForTests(t *testing.T) *Handlers {
	if handlersForTests == nil {

		confForTests, err := server_config.NewConfig()
		if err != nil {
			t.Log(err)
		}

		loggerForTests, err := mlg.NewZapLogger(confForTests.LogLevel)
		if err != nil {
			t.Log(err)
		}

		authForTests, err := myauth.Initialize(confForTests, loggerForTests)
		if err != nil {
			t.Log(err)
		}

		storeForTests, err := store.NewStorage(confForTests, loggerForTests)
		if err != nil {
			t.Log(err)
		}

		servicesForTests := services.NewDBServices(storeForTests, confForTests, loggerForTests, *authForTests)
		handlersForTests, err = NewHandlers(storeForTests, confForTests, loggerForTests, *authForTests, servicesForTests)
		if err != nil {
			t.Log(err)
		}
	}
	return handlersForTests
}

func TestHandlers_GetUserID(t *testing.T) {

	// сначала описываем структуру - скелет ответа, который для нас важен в тесте хэндлера.
	type want struct {
		userID int
		isAuth bool
		err    error
	}
	//далее описываем тест кейсы. Каждый тест кейс содержит имя и экземпляр структуры want с конкретными значениями
	tests := []struct {
		name string
		want want
	}{
		{
			name: "negative test #1", // первый кейс имеет имя positive test #1
			want: want{
				userID: 0,
				isAuth: false,
				err:    nil,
			},
		},
	}

	// перебираем все тест кейсы ( в нашем случае он всего один)
	for _, test := range tests { // значение будет храниться в переменной test
		t.Run(test.name, func(t *testing.T) { // запускаем в горутине каждый кейс параллельно
			handlersForTests := initHandlersForTests(t)
			request := httptest.NewRequest(http.MethodGet, "/api/user/get-data-items", nil)
			resultUserID, resultIsAuth, resultErr := handlersForTests.GetUserID(request)
			assert.Equal(t, test.want.userID, resultUserID)
			assert.Equal(t, test.want.isAuth, resultIsAuth)
			assert.Equal(t, test.want.err, resultErr)
		})
	}
}
