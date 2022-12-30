package main

import (
	"github.com/vasiliyantufev/go-advanced/internal/app"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	rtr := app.NewRouter()
	log.Printf("Starting application on port %v\n", portNumber)
	con := http.ListenAndServe(portNumber, rtr)
	if con != nil {
		log.Fatal(con)
	}
}
