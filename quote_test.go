
package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestQuoteAPI(t *testing.T) {
	LoadQuotes()
	defer SaveQuotes()

	router := setupRouter()

	// Test Create
	body := []byte(`{"author":"TestAuthor","quote":"Test quote"}`)
	req := httptest.NewRequest("POST", "/quotes", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusCreated || !strings.Contains(resp.Body.String(), "Цитата добавлена") {
		t.Errorf("POST /quotes failed: %v", resp.Body.String())
	}

	// Test Get all
	req = httptest.NewRequest("GET", "/quotes", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if !strings.Contains(resp.Body.String(), "TestAuthor") {
		t.Errorf("GET /quotes failed: %v", resp.Body.String())
	}

	// Test filter by author
	req = httptest.NewRequest("GET", "/quotes?author=TestAuthor", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if !strings.Contains(resp.Body.String(), "Test quote") {
		t.Errorf("GET /quotes?author failed: %v", resp.Body.String())
	}

	// Test random quote
	req = httptest.NewRequest("GET", "/quotes/random", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if !strings.Contains(resp.Body.String(), "Случайная цитата") {
		t.Errorf("GET /quotes/random failed: %v", resp.Body.String())
	}

	// Test delete
	req = httptest.NewRequest("DELETE", "/quotes/1", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if !strings.Contains(resp.Body.String(), "успешно удалена") {
		t.Errorf("DELETE /quotes/1 failed: %v", resp.Body.String())
	}
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/quotes", CreateQuote).Methods("POST")
	r.HandleFunc("/quotes", GetQuotes).Methods("GET")
	r.HandleFunc("/quotes/random", GetRandomQuote).Methods("GET")
	r.HandleFunc("/quotes/{id}", DeleteQuote).Methods("DELETE")
	return r
}
