package utils

import (
    "encoding/json"
    "net/http"
)

// RespondWithError sends an error response with the given status code and message
func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    errorResponse := map[string]string{"error": message}
    _ = json.NewEncoder(w).Encode(errorResponse)
}

// RespondWithJSON sends a JSON response with the given status code and data
func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    _ = json.NewEncoder(w).Encode(data)
}
