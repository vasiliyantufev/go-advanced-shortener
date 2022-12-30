package app

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/", PostHandler).Methods("POST")
	rtr.HandleFunc("/", IndexHandler)
	rtr.HandleFunc("/{id}", GetHandler).Methods("GET")

	return rtr
}
