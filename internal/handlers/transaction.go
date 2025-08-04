package handlers

import (
	"banking-ledger/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"encoding/json"

)

func (h *Handler) NewTransaction(c *gin.Context) {
    var req struct {
        AccountID string `json:"account_id"`
        Type      string `json:"type"`
        Amount    int64  `json:"amount"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    txn := models.Transaction{
        ID:        uuid.New().String(),
        AccountID: req.AccountID,
        Type:      req.Type,
        Amount:    req.Amount,
        Status:    "pending",
        CreatedAt: time.Now(),
    }

    body, err := json.Marshal(txn)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to encode transaction"})
        return
    }

    err = h.S.MQ.Publish(
        "", "transactions", false, false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to enqueue transaction"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":        "Transaction submitted",
        "transaction_id": txn.ID,
    })
}

func (h *Handler) GetTransactionHistory(c *gin.Context) {
	accountID := c.Param("id")
	txns, err := h.S.GetTransactionsByAccountID(accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch transactions"})
		return
	}
	if len(txns) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no transactions found"})
		return
	}
	c.JSON(http.StatusOK, txns)
}
