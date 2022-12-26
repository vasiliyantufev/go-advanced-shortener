package app

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

var data = NewDM()

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Index</h1>"))
}

func GetHandler(w http.ResponseWriter, r *http.Request) {

	short := mux.Vars(r)
	if short["id"] == "" {
		http.Error(w, "The query parameter is missing", http.StatusBadRequest)
		return
	}

	link, ok := data.Get(short["id"])
	if !ok {
		http.Error(w, "Link not found", http.StatusInternalServerError)
		return
	}

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
	data.Put(short, string(resp))

	link := "http://" + r.Host + "/" + short
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(link))
}
