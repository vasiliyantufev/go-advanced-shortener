package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const testURL = "https://rambler.com/"

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

	short := shorting()
	urls[short] = string(testURL)

	w := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/{id}", GetHandler).Methods("GET")
	//Greeter("Hello").AddRoute(r)
	r.ServeHTTP(w, httptest.NewRequest("GET", "http://127.0.0.1:8080/" + short, nil))

	loc := w.Header().Get("Location")
	status := w.Code

	assert.Equal(t, loc, testURL,
		fmt.Sprintf("Incorrect location. Expect %s, got %s", loc, testURL))

	assert.Equal(t, status, http.StatusTemporaryRedirect,
		fmt.Sprintf("Incorrect status. Expect %v, got %v", status, http.StatusInternalServerError))
}

func TestPostHandler(t *testing.T) {

	req, err := http.NewRequest("POST", "/", strings.NewReader(testURL))
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

