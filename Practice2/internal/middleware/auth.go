package middleware

import (
	"GolangLessons/Practice2/models"
	"encoding/json"
	"log"
	"net/http"
)

func APIKeyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)

		apiKey := r.Header.Get("X-API-Key")

		if apiKey != "secret123" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "unauthorized"})
			return
		}

		next(w, r)
	}
}
