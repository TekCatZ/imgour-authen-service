package ping

import (
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/usecases/handler/ping"
	"github.com/gin-gonic/gin"
)

func Ping(router *gin.RouterGroup) {
	router.GET("/ping", ping.PingHandle)
}
