package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/timut2/avito_test_task/internal/middleware"
	"github.com/timut2/avito_test_task/internal/models"
	"github.com/timut2/avito_test_task/internal/service"
	"github.com/timut2/avito_test_task/internal/utils"
)

type sendHandler struct {
	sendService *service.SendService
}

func NewSendHandler(sendService *service.SendService) *sendHandler {
	return &sendHandler{sendService: sendService}
}

func (h *sendHandler) Send(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request method")
		return
	}
	var req models.SendCoinRequest
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		utils.SendErrorResponse(w, http.StatusUnauthorized, "User ID not found in context")
	}

	log.Println("Validated UserID:", userID)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if req.Amount <= 0 {
		utils.SendErrorResponse(w, http.StatusBadRequest, "неправильный формат для количистава перевода средст")
	}

	err := h.sendService.Send(userID, req)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode("Transaction went successfully")
	if err != nil {
		log.Printf("Error with decoding JSON forresponse: %v", err)
	}
}
