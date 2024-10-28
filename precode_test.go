package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {

	totalCount := len(cafeList["moscow"])
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//Первая проверка
	//проверяю код запроса 200
	require.Equal(t, responseRecorder.Code, http.StatusOK)
	//проверяю тело запроса на заполненность
	assert.NotEmpty(t, responseRecorder.Body)

	//Вторая проверка
	//Проверяю на правильность заполненеия города
	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)

	//Третья проверка
	//получаю параметр count
	count := req.URL.Query().Get("count")
	//сравниваю totalCount и count
	assert.Equal(t, totalCount, count)
}
