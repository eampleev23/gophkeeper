package handlers

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	url2 "net/url"
	"testing"
)

func TestHandlers_AddText(t *testing.T) {

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
				code: http.StatusUnauthorized, // код ответа хэндлера 200
			},
		},
	}
	// перебираем все тест кейсы ( в нашем случае он всего один)
	for _, test := range tests { // значение будет храниться в переменной test

		t.Run(test.name, func(t *testing.T) { // запускаем в горутине каждый кейс параллельно

			var textDataNewItemRequestStr = `{"meta-value": "`
			textDataNewItemRequestStr += "фейковые данные для автотеста"
			textDataNewItemRequestStr += `", "text-content": "`
			textDataNewItemRequestStr += "Мороз и солнце; день чудесный! Еще ты дремлешь"
			textDataNewItemRequestStr += `,"}`

			var textDataNewItemRequestBytes = []byte(textDataNewItemRequestStr)

			handler := "/api/user/add-text"
			url, err := url2.JoinPath("http://localhost:8080", handler)
			if err != nil {
				t.Fatal(err)
			}

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(textDataNewItemRequestBytes))
			if err != nil {
				t.Fatal(err)
			}

			request.Header.Set("Content-Type", "application/json; charset=UTF-8")
			client := &http.Client{}

			response, err := client.Do(request)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, test.want.code, response.StatusCode)
		})
	}
}
