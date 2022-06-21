package route

import (
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/configs"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/controllers/route/authen"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/controllers/route/ping"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/usecases/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(serverConfig configs.ServerConfig) *http.Server {
	router := gin.Default()

	router.Use(middleware.CorsMiddleware(serverConfig))
	router.Use(middleware.AuthMiddleware)

	apiV1 := router.Group("/api/v1", middleware.VerifySession(nil))
	{
		authen.Setup(apiV1)
	}

	ping.Ping(router.Group(""))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return srv
}
