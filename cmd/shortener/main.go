package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"math/rand"
	"net/http"
)

const portNumber = ":8080"
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var urls = make(map[string]string)

// declaring a struct
type Url struct {
	// defining struct variables
	URL string
}

func shorting() string {
	b := make([]byte, 5)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Index</h1>"))
}

func main() {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/", PostHandler).Methods("POST")
	rtr.HandleFunc("/", index)
	rtr.HandleFunc("/{id}", GetHandler).Methods("GET")
	////rtr.HandleFunc("/{id:[0-9]+}", GetHandler).Methods("GET")

	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, rtr)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {

	short := mux.Vars(r)

	if short["id"] == "" {
		http.Error(w, "The query parameter is missing", http.StatusBadRequest)
		return
	}

	link := urls[short["id"]]

	w.WriteHeader(http.StatusTemporaryRedirect)
	w.Write([]byte(link))

	//http.Redirect(w, r, link, http.StatusTemporaryRedirect)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	//var urls map[string]string

	resp, err := ioutil.ReadAll(r.Body)
	var url Url

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Unmarshal(resp, &url)

	short := shorting()

	//urls["kkk"] = "jjj"

	urls[short] = url.URL

	//rp, err :=
	//urls = append(urls, url.URL)

	//urls = append(urls, string(url))

	//par := strings.Trim(url.URL, "\n")

	// устанавливаем заголовок Content-Type
	// для передачи клиенту информации, кодированной в JSON
	//w.Header().Set("content-type", "application/json")
	// устанавливаем статус-код 200
	w.WriteHeader(http.StatusCreated)
	//w.Write([]byte(par))

	//w.Write([]byte(url.URL))
	w.Write([]byte(short))

	//пишем тело ответа
	//w.Write([]byte(name.URL))
	//fmt.Fprintln(w, par)

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
