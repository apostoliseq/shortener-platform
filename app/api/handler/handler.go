package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/apostoliseq/shortener-platform/api/store"
)

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func Shorten(w http.ResponseWriter, r *http.Request) {
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

	record := store.Save(body.URL)
	json.NewEncoder(w).Encode(map[string]string{"code": record.Code})
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	record, ok := store.GetByCode(code)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, record.LongURL, http.StatusFound)
}
