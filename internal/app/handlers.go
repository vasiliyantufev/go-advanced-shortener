package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

var urls = make(map[string]string)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Index</h1>"))
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

	fmt.Print(r)


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
