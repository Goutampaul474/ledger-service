package handlers_test

import (
	"banking-ledger/internal/handlers"
	"banking-ledger/internal/models"
	"banking-ledger/internal/services"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// StubService embeds real Service but overrides CreateAccount
type StubService struct {
	services.Service
}

func (s *StubService) CreateAccount(name string, balance int64) (*models.Account, error) {
	return &models.Account{
		ID:        "fake-id",
		Name:      name,
		Balance:   balance,
		CreatedAt: time.Now(),
	}, nil
}

func TestCreateAccountHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// use stub service
	h := &handlers.Handler{S: &StubService{}}
	router := gin.New()
	router.POST("/accounts", h.CreateAccount)

	// simulate HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/accounts", strings.NewReader(`{"name":"Alice","balance":500}`))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	// assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var acc models.Account
	_ = json.Unmarshal(w.Body.Bytes(), &acc)
	assert.Equal(t, "Alice", acc.Name)
	assert.Equal(t, int64(500), acc.Balance)
}
