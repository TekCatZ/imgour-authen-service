package route

import (
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/configs"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/controllers/route/authen"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/usecases/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(serverConfig configs.ServerConfig) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CorsMiddleware(serverConfig))
	router.Use(middleware.AuthMiddleware)

	apiV1 := router.Group("/api/v1", middleware.VerifySession(nil))
	{
		authen.Setup(apiV1)
	}

	return router
}
