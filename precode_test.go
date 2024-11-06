package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тестирую статус ответа 200
func TestMainHandlerStausOk(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	//проверяю код запроса 200
	require.Equal(t, responseRecorder.Code, http.StatusOK, "Wrong response code")
	require.NotEmpty(t, responseRecorder.Body, "Empty response body")
}

// Тестирую правильность заполнения города
func TestMainHandlerCityOk(t *testing.T) {
	//Запрос другого города
	req := httptest.NewRequest("GET", "/cafe?count=10&city=spb", nil) // здесь нужно создать запрос к сервису
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	//проверяю код запроса 400
	require.Equal(t, responseRecorder.Code, http.StatusBadRequest, "Wrong response code")
	require.NotEmpty(t, responseRecorder.Body, "Wrong city value")
}

// Тестирую параметр count
func TestMainHandlerCount(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	//проверяю код запроса 200
	require.Equal(t, http.StatusOK, responseRecorder.Code, "Wrong response code")
	require.NotEmpty(t, responseRecorder.Body, "Empty response body")
	//сравниваю количество кафе
	quantity := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, quantity, 4, "Wrong cafe quantity")
}
