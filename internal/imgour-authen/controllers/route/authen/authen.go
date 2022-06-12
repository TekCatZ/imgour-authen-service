package authen

import (
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	router.Group("/auth")
	{
	}
}
