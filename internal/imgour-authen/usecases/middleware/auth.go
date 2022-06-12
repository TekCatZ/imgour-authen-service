package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/supertokens"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	supertokens.Middleware(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		c.Next()
	})).ServeHTTP(c.Writer, c.Request)
	// we call Abort so that the next handler in the chain is not called, unless we call Next explicitly
	c.Abort()
}
