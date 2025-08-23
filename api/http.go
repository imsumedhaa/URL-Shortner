package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/imsumedhaa/In-memory-database/pkg/client/postgres"
	shortner "github.com/imsumedhaa/URL-Shortner/Shortner"
)

type shortenResponse struct {
	ShortURL string `json:"Code"`
}

type ShortenRequest struct {
	OriginalURL string `json:"OgURL"`
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

func (h *Http) ShortenHandeler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid body request: %s", err), http.StatusBadRequest)
		return
	}

	if req.OriginalURL == "" {
		http.Error(w, "URL cannot be empty", http.StatusBadRequest)
		return
	}

	code := shortner.Generator(req.OriginalURL)

	shortURL := fmt.Sprint(w, "http://localhost:8080/%s", code)

	if err := h.client.CreatePostgresRow(req.OriginalURL, shortURL); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create row: %s", err), http.StatusInternalServerError)
		return
	}

	response := shortenResponse{ShortURL: shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
