package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestHandle(c *gin.Context) {
	c.JSON(http.StatusOK, "Ok")
}
