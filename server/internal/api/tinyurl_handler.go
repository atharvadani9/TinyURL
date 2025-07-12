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

type getTinyURLRequest struct {
	ShortURL string `json:"shortUrl"`
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

	existingTinyURL, err := h.tinyURLStore.GetByOriginalURL(req.OriginalURL)
	if err == nil && existingTinyURL != nil {
		h.logger.Printf("Returning existing tiny URL: %v", existingTinyURL.ShortURL)
		utils.WriteJSON(w, http.StatusOK, utils.Envelope{"shortUrl": existingTinyURL.ShortURL})
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

	h.logger.Printf("Tiny URL created: %v", tinyURL.ShortURL)
	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"shortUrl": tinyURL.ShortURL})
}

func generateShortURL() string {
	return utils.GenerateRandomString(6)
}

func (h *TinyURLHandler) GetTinyURL(w http.ResponseWriter, r *http.Request) {
	var req getTinyURLRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.Printf("Error decoding request body: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid request body"})
		return
	}

	tinyURL, err := h.tinyURLStore.Get(req.ShortURL)
	if err != nil {
		h.logger.Printf("Error getting tiny URL: %v", err)
		utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "Tiny URL not found"})
		return
	}

	h.logger.Printf("URL found: %v", tinyURL.OriginalURL)
	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"originalUrl": tinyURL.OriginalURL})
}
