package main

import (
	"fmt"
	"net/http"

	"github.com/apostoliseq/shortener-platform/api/handler"
)

func main() {
	http.HandleFunc("/health", handler.Health)
	http.HandleFunc("/shorten", handler.Shorten)
	http.HandleFunc("/", handler.Redirect)

	fmt.Println("API listening on :8080")
	http.ListenAndServe(":8080", nil)
}
