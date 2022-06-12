package middleware

import (
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/configs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func CorsMiddleware(serverConfig configs.ServerConfig) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{serverConfig.ServiceBaseUrl},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: append([]string{"content-type"},
			supertokens.GetAllCORSHeaders()...),
		AllowCredentials: true,
	})
}
