package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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

}

func TestPostHandler(t *testing.T) {

	url := "https://rambler.com/"

	req, err := http.NewRequest("POST", "/", strings.NewReader(url))
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


//tests := []struct {
//	name  string
//	value string
//	want  string
//}{
//	{name: "Correct", value: "<h1>Index</h1>", want: "<h1>Index</h1>"},
//}
//
//
//
//for _, tt := range tests {
//	t.Run(tt.name, func(t *testing.T) {
//		assert.Equal(t, tt.value, tt.want,
//			fmt.Sprintf("Incorrect result. Expect %s, got %s", tt.want, tt.value))
//	})
//}
