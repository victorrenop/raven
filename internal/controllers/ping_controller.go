package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthChecker is a simple controller for pinging
type HealthChecker struct{}

// PingHandler is the function responsible for health checking the app
func (h HealthChecker) PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
