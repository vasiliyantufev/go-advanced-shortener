package app

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const testURL = "https://rambler.com/"

func TestIndexHandler(t *testing.T) {

	w := httptest.NewRecorder()
	r := NewRouter()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/", nil))

	body := w.Body.String()
	status := w.Code

	expected := "<h1>Index</h1>"
	assert.Equal(t, body, expected,
		fmt.Sprintf("Incorrect body. Expect %s, got %s", body, expected))

	assert.Equal(t, status, http.StatusOK,
		fmt.Sprintf("Incorrect status. Expect %v, got %v", status, http.StatusOK))
}

func TestPostHandler(t *testing.T) {

	w := httptest.NewRecorder()
	r := NewRouter()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/", strings.NewReader(testURL)))

	status := w.Code

	assert.Equal(t, status, http.StatusCreated,
		fmt.Sprintf("Incorrect status. Expect %v, got %v", status, http.StatusCreated))
}

func TestGetHandler(t *testing.T) {

	short := shorting(testURL)
	data.Put(short, string(testURL))

	w := httptest.NewRecorder()
	r := NewRouter()
	r.ServeHTTP(w, httptest.NewRequest("GET", "http://127.0.0.1:8080/"+short, nil))

	loc := w.Header().Get("Location")
	status := w.Code

	assert.Equal(t, loc, testURL,
		fmt.Sprintf("Incorrect location. Expect %s, got %s", loc, testURL))

	assert.Equal(t, status, http.StatusTemporaryRedirect,
		fmt.Sprintf("Incorrect status. Expect %v, got %v", status, http.StatusInternalServerError))
}
