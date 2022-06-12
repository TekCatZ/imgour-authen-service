package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandle(c *gin.Context) {
	c.JSON(http.StatusOK, "Pong")
}
