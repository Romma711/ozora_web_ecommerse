package utils

import "net/http"

func UnauthorizedUser(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"User not authorized\"}"))
}