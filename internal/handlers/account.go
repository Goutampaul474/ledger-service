package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func (h *Handler) CreateAccount(c *gin.Context) {
	var req struct {
		Name    string `json:"name"`
		Balance int64  `json:"balance"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	acc, err := h.S.CreateAccount(req.Name, req.Balance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create account"})
		return
	}

	c.JSON(http.StatusOK, acc)
}

func (h *Handler) GetAccount(c *gin.Context) {
	id := c.Param("id")
	acc, err := h.S.GetAccountByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}
	c.JSON(http.StatusOK, acc)
}
