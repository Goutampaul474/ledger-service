package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HealthCheck(c *gin.Context) {
    healthy := h.S.IsPostgresHealthy() && h.S.IsMongoHealthy() && h.S.IsRabbitMQHealthy()
    if !healthy {
        c.JSON(http.StatusInternalServerError, gin.H{"status": "unhealthy"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

