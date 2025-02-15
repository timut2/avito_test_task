package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timut2/avito_test_task/internal/handlers"
	"github.com/timut2/avito_test_task/internal/middleware"
	"github.com/timut2/avito_test_task/internal/models"
	"github.com/timut2/avito_test_task/internal/repository/fake"
	"github.com/timut2/avito_test_task/internal/service"
)

func TestSend(t *testing.T) {

	sendRepo := fake.NewFakeSendRepo()
	sendService := service.NewSendService(sendRepo)
	h := handlers.NewSendHandler(sendService)

	w := httptest.NewRecorder()

	reqBody := models.SendCoinRequest{
		ToUser: "Bob",
		Amount: 100,
	}

	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/sendCoin", bytes.NewBuffer(jsonBody))
	ctx := context.WithValue(req.Context(), middleware.UserIDKey, 1)
	req = req.WithContext(ctx)

	h.Send(w, req)

	res := w.Result()
	defer res.Body.Close()

	senderBalance, recieverBalance := sendRepo.TransactionInfo(1)

	assert.Equal(t, senderBalance+200, recieverBalance)
}
