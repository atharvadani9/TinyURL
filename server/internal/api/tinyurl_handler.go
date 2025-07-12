package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"tinyurl/internal/store"
	"tinyurl/internal/utils"
)

type createTinyURLRequest struct {
	OriginalURL string `json:"originalUrl"`
}

type TinyURLHandler struct {
	tinyURLStore store.TinyURLStore
	logger       *log.Logger
}

func NewTinyURLHandler(tinyURLStore store.TinyURLStore, logger *log.Logger) *TinyURLHandler {
	return &TinyURLHandler{tinyURLStore: tinyURLStore, logger: logger}
}

func (h *TinyURLHandler) CreateTinyURL(w http.ResponseWriter, r *http.Request) {
	var req createTinyURLRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.Printf("Error decoding request body: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid request body"})
		return
	}

	if req.OriginalURL == "" {
		h.logger.Printf("Error: original URL is empty")
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Original URL is required"})
		return
	}

	shortURL := generateShortURL()
	for {
		_, err := h.tinyURLStore.Get(shortURL)
		if err == sql.ErrNoRows {
			break
		}
		shortURL = generateShortURL()
	}

	tinyURL := &store.TinyURL{
		OriginalURL: req.OriginalURL,
		ShortURL:    shortURL,
	}

	err = h.tinyURLStore.Create(tinyURL)
	if err != nil {
		h.logger.Printf("Error creating tiny URL: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Failed to create tiny URL"})
		return
	}

	h.logger.Printf("Tiny URL created: %v", tinyURL)
	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"tinyUrl": tinyURL})
}

func generateShortURL() string {
	return utils.GenerateRandomString(6)
}
