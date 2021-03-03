package registry

import (
	xapp "github.com/coder2m/component"
	"github.com/coder2m/reminder/internal/app/api/v1/middleware"
	R "github.com/coder2m/reminder/pkg/response"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	router *gin.Engine
)

func Engine() *gin.Engine {
	if router == nil {
		if xapp.Debug() {
			gin.DisableConsoleColor()
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}
		router = gin.New()
		router.NoRoute(R.HandleNotFound)
		router.Use(
			middleware.RecoverMiddleware(20*time.Second),
			middleware.XMonitor(),
			middleware.XTrace(),
		)
	}
	return router
}

