package utils

import (
	"net/http"

	"github.com/rs/cors"
)

func UnauthorizedUser(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"User not authorized\"}"))
}
func EnableCORS () *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	return c
}