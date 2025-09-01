package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/imsumedhaa/In-memory-database/pkg/client/postgres"
	shortner "github.com/imsumedhaa/URL-Shortner/Shortner"
	_ "github.com/lib/pq"
)

type Response struct {
	ShortURL    string `json:"ShortURL"`
	OriginalURL string `json:"OriginalURL"`
}

type Request struct {
	OriginalURL string `json:"OriginalURL"`
	ShortURL    string `json:"ShortURL"`
}

type Http struct {
	client postgres.Client
}

func NewHttp(host, port, username, password, dbname string) (*Http, error) {
	dbClient, err := postgres.NewClient(host, port, username, password, dbname)

	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}
	return &Http{client: dbClient}, nil
}

func (h *Http) Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid body request: %s", err), http.StatusBadRequest)
		return
	}

	if req.OriginalURL == "" {
		http.Error(w, "URL cannot be empty", http.StatusBadRequest)
		return
	}

	code := shortner.Generator(req.OriginalURL)

	shortURL := fmt.Sprintf("http://localhost:8080/%s", code)

	if err := h.client.CreatePostgresRow(req.OriginalURL, shortURL); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create row: %s", err), http.StatusInternalServerError)
		return
	}

	response := Response{ShortURL: shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func (h *Http) GetOriginal(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	short := r.URL.Query().Get("short")
	if short == "" {
		http.Error(w, "Short url cannot be empty", http.StatusBadRequest)
		return
	}

	ogUrl, err := h.client.GetPostgresRow(short)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get the row: %s", err), http.StatusInternalServerError)
		return
	}

	response := Response{OriginalURL: ogUrl}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func (h *Http) DeleteShortUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid json body", http.StatusBadRequest)
		return
	}

	if req.OriginalURL == "" {
		http.Error(w, "Original Url cannot be empty", http.StatusBadRequest)
		return
	}

	if err := h.client.DeletePostgresRow(req.OriginalURL); err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete the row: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Short URL for '%s' deleted successfully", req.OriginalURL),
	})

}

func (h *Http) Redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := strings.TrimPrefix(r.URL.Path, "/")
	if shortCode == "" {
		http.NotFound(w, r)
		return
	}

	originalURL, err := h.client.GetPostgresRow(shortCode)
	if err != nil || originalURL == "" {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func (h *Http) Run() error {
	h.routes()
	log.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}
	return nil
}

func (h *Http) routes() {
	http.HandleFunc("/create", h.Shorten)
	http.HandleFunc("/get", h.GetOriginal)
	http.HandleFunc("/delete", h.DeleteShortUrl)
	http.HandleFunc("/", h.Redirect)

}
