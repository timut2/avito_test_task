package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/timut2/avito_test_task/internal/models"
)

func SendErrorResponse(w http.ResponseWriter, statusCode int, errorMessage string) {
	log.Printf("Error: %s, Status code: %d", errorMessage, statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := models.ErrorResponse{
		Errors: errorMessage,
	}

	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		log.Printf("Error with decoding JSON forresponse: %v", err)
	}
}
