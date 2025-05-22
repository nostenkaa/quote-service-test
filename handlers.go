package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateQuote(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Author string `json:"author"`
		Quote  string `json:"quote"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Некорректный формат входных данных", http.StatusBadRequest)
		return
	}

	q := AddQuote(input.Author, input.Quote)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, " Цитата добавлена:\nID: %d\nАвтор: %s\nТекст: %s\n", q.ID, q.Author, q.Quote)
}

func GetQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	var result []Quote
	if author != "" {
		result = GetQuotesByAuthor(author)
	} else {
		result = GetAllQuotes()
	}

	var sb strings.Builder
	for _, q := range result {
		sb.WriteString(fmt.Sprintf("ID: %d\nАвтор: %s\nЦитата: %s\n\n", q.ID, q.Author, q.Quote))
	}
	if sb.Len() == 0 {
		sb.WriteString("Цитаты не найдены\n")
	}
	w.Write([]byte(sb.String()))
}

func GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	if len(quotes) == 0 {
		http.Error(w, "Нет доступных цитат", http.StatusNotFound)
		return
	}
	q := quotes[rand.Intn(len(quotes))]
	fmt.Fprintf(w, " Случайная цитата:\nID: %d\nАвтор: %s\nЦитата: %s\n", q.ID, q.Author, q.Quote)
}

func DeleteQuote(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}
	if ok := DeleteQuoteByID(id); !ok {
		http.Error(w, "Цитата не найдена", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "🗑Цитата с ID %d успешно удалена\n", id)
}
