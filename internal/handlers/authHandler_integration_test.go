package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/timut2/avito_test_task/internal/handlers"
	"github.com/timut2/avito_test_task/internal/models"
	"github.com/timut2/avito_test_task/internal/repository"
	"github.com/timut2/avito_test_task/internal/repository/fake"
	repoUserMock "github.com/timut2/avito_test_task/internal/repository/mocks"
	"github.com/timut2/avito_test_task/internal/service"
	"github.com/timut2/avito_test_task/internal/utils"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repoUserMock.NewMockUserRepo(ctrl)

	reqBody := models.AuthRequest{
		Username: "testuser",
		Password: "Password123!",
	}

	mockRepo.EXPECT().FindByUserName(reqBody.Username).Return(nil, repository.ErrUserNotFound).Times(1)
	mockRepo.EXPECT().Create(gomock.Any()).Return(1, nil).Times(1)

	authService := service.NewAuthService(mockRepo)
	h := handlers.NewAuthHandler(authService)

	w := httptest.NewRecorder()
	requestBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/auth", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	h.Login(w, req)

	res := w.Result()
	defer res.Body.Close()

	var response models.AuthResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)
	require.NotEmpty(t, response.Token)
	realToken, err := utils.GenerateJWT(1)
	require.NoError(t, err)
	require.Equal(t, realToken, response.Token)
}

func TestLoginWithFakeRepo(t *testing.T) {

	repo := fake.NewFakeUserRepo()

	authService := service.NewAuthService(repo)
	h := handlers.NewAuthHandler(authService)

	w := httptest.NewRecorder()

	reqBody := models.AuthRequest{
		Username: "testuser",
		Password: "Password123!",
	}

	requestBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/auth", bytes.NewBuffer(requestBody))
	h.Login(w, req)
	res := w.Result()
	defer res.Body.Close()

	var response models.AuthResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)
	require.NotEmpty(t, response.Token)
	realToken, err := utils.GenerateJWT(1)
	require.NoError(t, err)
	realId, err := utils.ValidateToken(realToken)
	require.NoError(t, err)
	Id, err := utils.ValidateToken(response.Token)
	require.NoError(t, err)
	require.Equal(t, realId, Id)
}
