package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vasiliyantufev/go-advanced/internal/app"
	"net/http"
)

const portNumber = ":8080"

func main() {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/", app.PostHandler).Methods("POST")
	rtr.HandleFunc("/", app.IndexHandler)
	rtr.HandleFunc("/{id}", app.GetHandler).Methods("GET")

	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, rtr)
}
