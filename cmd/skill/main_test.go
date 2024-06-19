package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebhook(t *testing.T) {
	// описывает ожидаемое тело ответа при успешном запросе
	sucsessBody := `{
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
			expectedBody: sucsessBody,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, "/", nil)
			w := httptest.NewRecorder()

			// вызовем хэндлер как обычную функцию, без запуска сервера
			webhook(w, r)

			assert.Equal(t, w.Code, tc.expectedCode, "Код ответа не совпадает с ожидаемым")
			// проверим корректность полученного тела ответа, если мы его ожидаем
			if tc.expectedBody != "" {
				// assert.JSONEq помогает сравнить две JSON-строки
				assert.JSONEq(t, w.Body.String(), tc.expectedBody, "Тело ответа не совпадает с ожидаемым")
			}
		})
	}
}
