package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"math/rand"
	"net/http"
)

const portNumber = ":8080"
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var urls = make(map[string]string)

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
	w.Header().Set("Location", link)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	short := shorting()
	urls[short] = string(resp)

	link := "http://" + r.Host + "/" + short
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(link))
}
