package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	urlStore = sync.Map{}
	letters  = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	baseURL  = "http://localhost:8080/r/"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateShortURL(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL(6)
	urlStore.Store(shortURL, req.URL)

	resp := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: baseURL + shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/r/"):]

	if originalURL, ok := urlStore.Load(shortURL); ok {
		http.Redirect(w, r, originalURL.(string), http.StatusFound)
	} else {
		http.Error(w, "URL not found", http.StatusNotFound)
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Server is running at port 8080"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", HealthCheckHandler)
	http.HandleFunc("/shorten", ShortenURLHandler)
	http.HandleFunc("/r/", RedirectHandler)
	fmt.Println("Server is running at port 8080")
	http.ListenAndServe(":8080", nil)
}
