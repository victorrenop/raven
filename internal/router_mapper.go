package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/victorrenop/raven/internal/controllers"
)

// WrappedRouter instance containing the gin gonic engine and the ping controller
type WrappedRouter struct {
	Router        *gin.Engine
	HealthChecker *controllers.HealthChecker
}

// Function that initiates all of the required controllers and endpoints
func (wrappedRouter *WrappedRouter) setupRouter() {
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

// Run function instatiates every controller, route and the API itself
func (wrappedRouter *WrappedRouter) Run(port string) error {
	wrappedRouter.setupRouter()

	return wrappedRouter.Router.Run(":" + port)
}
