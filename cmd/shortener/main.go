package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const portNumber = ":8080"

var urls []string

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Index</h1>"))
}

func main() {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/", PostHandler).Methods("POST")
	rtr.HandleFunc("/", index)
	rtr.HandleFunc("/{id:[0-9]+}", GetHandler)

	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, rtr)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	intVar, err := strconv.Atoi(vars["id"])
	// обрабатываем ошибку
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, urls[intVar], 307)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	url := r.FormValue("url")
	// обрабатываем ошибку
	if url == "" {
		http.Error(w, "The url parameter is missing", http.StatusBadRequest)
		return
	}

	urls = append(urls, string(url))
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(201)
	// пишем в тело ответа
	fmt.Fprintln(w, url)
}
