package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

const portNumber = ":8080"

var urls []string

// declaring a struct
type Url struct {
	// defining struct variables
	URL string
}

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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, urls[intVar], http.StatusTemporaryRedirect)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := ioutil.ReadAll(r.Body)
	var url Url

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Unmarshal(resp, &url)

	//rp, err :=
	urls = append(urls, url.URL)

	//urls = append(urls, string(url))

	// устанавливаем заголовок Content-Type
	// для передачи клиенту информации, кодированной в JSON
	//w.Header().Set("content-type", "application/json")
	// устанавливаем статус-код 200
	w.WriteHeader(http.StatusCreated)
	//пишем тело ответа
	//w.Write([]byte(name.URL))
	fmt.Fprintln(w, url.URL)

	//url := r.FormValue("URL")
	// обрабатываем ошибку
	//if url == "" {
	//	http.Error(w, "The url parameter is missing", http.StatusBadRequest)
	//	return
	//}
	//
	//urls = append(urls, string(url))
	//w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	//w.WriteHeader(201)
	//// пишем в тело ответа
	//fmt.Fprintln(w, url)
}
