package middleware

import "github.com/gin-gonic/gin"

func Test(c *gin.Context) {
	c.Request.Header.Add("test-mdw", "true")
	c.Next()
}
