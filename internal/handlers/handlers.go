package handlers

import (
	"banking-ledger/internal/models"
	"banking-ledger/internal/services"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	S *services.Service
}

func (h *Handler) CreateAccount(c *gin.Context) {
	var req struct {
		Name    string `json:"name"`
		Balance int64  `json:"balance"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	acc, err := h.S.CreateAccount(context.Background(), req.Name, req.Balance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, acc)
}

func (h *Handler) NewTransaction(c *gin.Context) {
	var req struct {
		AccountID string `json:"account_id"`
		Type      string `json:"type"` // deposit | withdraw
		Amount    int64  `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := models.Transaction{
		ID:        time.Now().Format("20060102150405"),
		AccountID: req.AccountID,
		Type:      req.Type,
		Amount:    req.Amount,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	if err := h.S.PublishTransaction(tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, tx)
}
