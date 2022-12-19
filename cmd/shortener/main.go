package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

const portNumber = ":8060"

var urls []string

// HelloWorld — обработчик запроса.
func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Index</h1>"))
}

func main() {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index)
	rtr.HandleFunc("/get_url", QueryHandler).Methods("GET")
	rtr.HandleFunc("/create", BodyHandler).Methods("POST")

	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, rtr)
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	// извлекаем фрагмент id= из URL запроса search?query=something
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "The id parameter is missing", http.StatusBadRequest)
		return
	}

	intVar, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(201)
	w.Write([]byte(urls[intVar]))
}

func BodyHandler(w http.ResponseWriter, r *http.Request) {
	// читаем Body
	url, err := io.ReadAll(r.Body)

	// обрабатываем ошибку
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	urls = append(urls, string(url))

	w.WriteHeader(307)
	w.Write([]byte(url))
}
