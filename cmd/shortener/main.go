package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const portNumber = ":8080"

var urls []string

// HelloWorld — обработчик запроса.
func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Index</h1>"))
}

func main() {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/", BodyHandler).Methods("POST")
	rtr.HandleFunc("/", index)
	rtr.HandleFunc("/get_url", QueryHandler).Methods("GET")

	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, rtr)
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	// извлекаем фрагмент id= из URL запроса search?query=something
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "The id parameter is missing", http.StatusBadRequest)
		return
	}

	intVar, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//w.WriteHeader(201)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// если пароль не верен, указываем код ошибки в заголовке
	w.WriteHeader(307)
	//fmt.Fprintln(w, []byte(urls[intVar]))
	http.Redirect(w, r, urls[intVar], http.StatusSeeOther)

	//w.Write([]byte(urls[intVar]))
	// пишем в тело ответа
	//fmt.Fprintln(w, url)
}

func BodyHandler(w http.ResponseWriter, r *http.Request) {

	// читаем Body
	//url, err := io.ReadAll(r.Body)
	//
	//// обрабатываем ошибку
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}

	url := r.FormValue("url")
	// обрабатываем ошибку
	if url == "" {
		http.Error(w, "The url parameter is missing", http.StatusBadRequest)
		return
	}

	urls = append(urls, string(url))
	//w.WriteHeader(307)
	//w.Write([]byte(url))
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(201)
	//w.Write([]byte(url))

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// если пароль не верен, указываем код ошибки в заголовке
	w.WriteHeader(201)
	// пишем в тело ответа
	fmt.Fprintln(w, url)
}
