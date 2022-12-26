package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const testUrl = "https://rambler.com/"

func TestIndexHandler(t *testing.T) {

	// Создаем запрос с указанием нашего хендлера. Нам не нужно
	// указывать параметры, поэтому вторым аргументом передаем nil
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Мы создаем ResponseRecorder(реализует интерфейс http.ResponseWriter)
	// и используем его для получения ответа
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)

	// Наш хендлер соответствует интерфейсу http.Handler, а значит
	// мы можем использовать ServeHTTP и напрямую указать
	// Request и ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Проверяем код
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Проверяем тело ответа
	expected := "<h1>Index</h1>"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetHandler(t *testing.T) {

	//short := shorting()
	//urls[short] = string(testUrl)
	//
	//req, err := http.NewRequest("GET", "http://127.0.0.1:8080/" + short, nil)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(GetHandler)
	//handler.ServeHTTP(rr, req)
	//
	//if status := rr.Code; status != http.StatusTemporaryRedirect {
	//	t.Errorf("handler returned wrong status code: got %v want %v",
	//		status, http.StatusCreated)
	//}
}

func TestPostHandler(t *testing.T) {

	req, err := http.NewRequest("POST", "/", strings.NewReader(testUrl))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

