package handlers_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/timut2/avito_test_task/internal/handlers"
	"github.com/timut2/avito_test_task/internal/middleware"
	"github.com/timut2/avito_test_task/internal/repository/fake"
	"github.com/timut2/avito_test_task/internal/service"
)

func TestBuy(t *testing.T) {
	purchaseRepo := fake.NewFakePurchaseRepo()
	buyService := service.NewBuyService(purchaseRepo)
	h := handlers.NewBuyHandler(buyService)

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/api/buy/1", bytes.NewBuffer([]byte{}))
	ctx := context.WithValue(req.Context(), middleware.UserIDKey, 1)
	req = req.WithContext(ctx)

	h.Buy(w, req)

	res := w.Result()
	defer res.Body.Close()

	// Validate purchase record
	require.Equal(t, 1, len(purchaseRepo.PurchaseRepo))
	require.Equal(t, purchaseRepo.PurchaseRepo[1], 1)

	// Validate balance update
	balance, itemCost := purchaseRepo.GiveInfo(1, 1)
	require.Equal(t, balance+itemCost, 1000)
}
