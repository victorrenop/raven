package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthChecker struct{}

func (h HealthChecker) PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
