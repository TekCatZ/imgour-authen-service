package profile

import (
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/usecases/handler/profile"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	router.Group("/profile")
	{
		router.GET("/me", profile.GetProfileHandler)
	}
}
