package authen

import (
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/usecases/handler/auth"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/usecases/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	router.Group("/auth", middleware.Test)
	{
		router.GET("/test", auth.TestHandle)
	}
}
