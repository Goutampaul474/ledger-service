package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"postgres": h.S.PG != nil,
		"mongodb":  h.S.MongoDB != nil,
		"rabbitmq": h.S.MQ != nil,
	})
}
