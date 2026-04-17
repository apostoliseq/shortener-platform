package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// URLRecord represents a stored URL mapping
type URLRecord struct {
	ID        int
	Code      string
	LongURL   string
	CreatedAt time.Time
}

// in-memory store: code → URLRecord
var store = map[string]URLRecord{}
var nextID = 1

func main() {
	// Register a handler to a path
	http.HandleFunc("/health", healthHandler) // register path → handler in DefaultServeMux
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	// Start listening on port 8080
	fmt.Println("API listening on :8080")
	http.ListenAndServe(":8080", nil) // listen on 8080, use DefaultServeMux to route
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body struct {
		URL string `json:"url"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	if body.URL == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}

	code := generateCode(6)
	store[code] = URLRecord{
		ID:        nextID,
		Code:      code,
		LongURL:   body.URL,
		CreatedAt: time.Now(),
	}
	nextID++

	json.NewEncoder(w).Encode(map[string]string{"code": code})
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:] // strip the leading "/"
	record, ok := store[code]
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, record.LongURL, http.StatusFound)
}

func generateCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[r.Intn(len(charset))]
	}
	return string(code)
}

// Define what happens when the path is hit
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok") // write "ok" back to the caller
}
