package utils

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Envelope map[string]any

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func WriteJSON(w http.ResponseWriter, status int, data Envelope) error {
	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func GenerateRandomString(length int) string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var result strings.Builder
	result.Grow(length)

	for i := 0; i < length; i++ {
		result.WriteByte(base62Chars[rand.Intn(len(base62Chars))])
	}

	return result.String()
}
