package main

import (
	"fmt"
	"github.com/vasiliyantufev/go-advanced/internal/app"
	"net/http"
)

const portNumber = ":8080"

func main() {

	rtr := app.NewRouter()

	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, rtr)
}