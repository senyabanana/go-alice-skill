package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebhook(t *testing.T) {
	// тип http.HandlerFunc реализует интерфейс http.Handler
	// это позволит передать хэндлер тестовому серверу
	handler := http.HandlerFunc(webhook)
	// запускаем тестовый сервер, будет выбран первый свободный порт
	srv := httptest.NewServer(handler)
	// останавливаем сервер после завершения теста
	defer srv.Close()

	// описывает ожидаемое тело ответа при успешном запросе
	successBody := `{
		"response": {
			"text": "Извините, я пока ничего не умею"
		},
		"version": "1.0"
	}`

	// описываем набор данных: метод запроса, ожидаемый код ответа, ожидаемое тело
	testCases := []struct {
		method       string
		expectedCode int
		expectedBody string
	}{
		{
			method:       http.MethodGet,
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "",
		},
		{
			method:       http.MethodPut,
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "",
		},
		{
			method:       http.MethodDelete,
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "",
		},
		{
			method:       http.MethodPost,
			expectedCode: http.StatusOK,
			expectedBody: successBody,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			//r := httptest.NewRequest(tc.method, "/", nil)
			//w := httptest.NewRecorder()
			//
			//// вызовем хэндлер как обычную функцию, без запуска сервера
			//webhook(w, r)

			//assert.Equal(t, w.Code, tc.expectedCode, "Код ответа не совпадает с ожидаемым")
			// проверим корректность полученного тела ответа, если мы его ожидаем
			//if tc.expectedBody != "" {
			//	// assert.JSONEq помогает сравнить две JSON-строки
			//	assert.JSONEq(t, w.Body.String(), tc.expectedBody, "Тело ответа не совпадает с ожидаемым")
			//}

			// делаем запрос с помощью библиотеки resty к адресу запущенного сервера,
			// который хранится в поле URL соответствующей структуры
			req := resty.New().R()
			req.Method = tc.method
			req.URL = srv.URL

			resp, err := req.Send()
			assert.NoError(t, err, "error making HTTP request")

			assert.Equal(t, tc.expectedCode, resp.StatusCode(), "Response code didn't match expected")
			// проверяем корректность полученного тела ответа, если мы его ожидаем
			if tc.expectedBody != "" {
				assert.JSONEq(t, tc.expectedBody, string(resp.Body()))
			}
		})
	}
}
