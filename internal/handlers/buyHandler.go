package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/timut2/avito_test_task/internal/middleware"
	"github.com/timut2/avito_test_task/internal/service"
	"github.com/timut2/avito_test_task/internal/utils"
)

type BuyHandler struct {
	BuyService *service.BuyService
}

func NewBuyHandler(buyService *service.BuyService) *BuyHandler {
	return &BuyHandler{BuyService: buyService}
}

func (h *BuyHandler) Buy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendErrorResponse(w, http.StatusMethodNotAllowed, "Неверный метод реквеста")
		return
	}

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		utils.SendErrorResponse(w, http.StatusUnauthorized, "Не удалось получить User ID из контекста")
		return
	}

	itemIDStr := strings.TrimPrefix(r.URL.Path, "/api/buy/")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil || itemID <= 0 {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Неправильный URL параметр")
		return
	}

	log.Printf("User %d пытается купить товар с ID: %d", userID, itemID)

	err = h.BuyService.BuyItem(userID, itemID)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode("Purchase successful")
	if err != nil {
		log.Printf("Error with decoding JSON forresponse: %v", err)
	}
}
