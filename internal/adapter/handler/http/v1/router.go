package http

import (
	"go-hex-arch/internal/adapter/handler/http/v1/middleware"
	"go-hex-arch/internal/configs"
	"go-hex-arch/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Router is a wrapper for HTTP router
type Router struct {
	*gin.Engine
}

func NewRouter(
	config *configs.App,
	log *logger.Logger,
	managerAccountHandler *ManagerAccountHandler,

) *Router {
	// Disable debug mode in production
	if !config.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	// CORS
	// ginConfig := cors.DefaultConfig()
	// allowedOrigins := config.AllowedOrigins
	// originsList := strings.Split(allowedOrigins, ",")
	// ginConfig.AllowOrigins = originsList

	router := gin.New()
	router.Use(middleware.Logger(log))
	// router.Use(sloggin.New(slog.Default()), gin.Recovery(), cors.New(ginConfig))

	v1 := router.Group("/v1")
	{

		category := v1.Group("/categories")
		{

			category.GET("/:id", managerAccountHandler.Get)
			category.GET("", managerAccountHandler.GetList)
			category.POST("", managerAccountHandler.Create)

		}

	}

	return &Router{
		router,
	}
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
