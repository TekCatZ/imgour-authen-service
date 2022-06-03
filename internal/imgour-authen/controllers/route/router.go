package route

import (
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/controllers/route/authen"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()
	apiV1 := router.Group("/api/v1")
	{
		authen.Setup(apiV1)
	}

	return router
}
