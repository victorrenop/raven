package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/victorrenop/raven/internal/controllers"
)

type WrappedRouter struct {
	Router        *gin.Engine
	HealthChecker *controllers.HealthChecker
}

func (wrappedRouter *WrappedRouter) SetupRouter() {
	if wrappedRouter.Router == nil {
		wrappedRouter.Router = gin.Default()
	}
	healthChecker := &controllers.HealthChecker{}
	wrappedRouter.HealthChecker = healthChecker
	wrappedRouter.mapRoutes()
}

func (wrappedRouter *WrappedRouter) mapRoutes() {
	router := wrappedRouter.Router
	router.GET("/ping", wrappedRouter.HealthChecker.PingHandler)
}

func (wrappedRouter *WrappedRouter) Run(port string) error {
	wrappedRouter.SetupRouter()

	return wrappedRouter.Router.Run(":" + port)
}
