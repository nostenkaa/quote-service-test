package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	LoadQuotes()
	defer SaveQuotes()

	router := mux.NewRouter()
	router.HandleFunc("/quotes", CreateQuote).Methods("POST")
	router.HandleFunc("/quotes", GetQuotes).Methods("GET")
	router.HandleFunc("/quotes/random", GetRandomQuote).Methods("GET")
	router.HandleFunc("/quotes/{id}", DeleteQuote).Methods("DELETE")

	fmt.Println(" Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
