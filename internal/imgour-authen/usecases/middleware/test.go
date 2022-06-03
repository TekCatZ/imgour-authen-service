package middleware

import "github.com/gin-gonic/gin"

func Test(c *gin.Context) {
	c.Header("test-mdw", "true")
}
