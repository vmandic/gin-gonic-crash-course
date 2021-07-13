package middlewares

import "github.com/gin-gonic/gin"

func BasicAuth() gin.HandlerFunc {
	// NOTE: takes a map with username / pwd pairs
	return gin.BasicAuth(gin.Accounts{
		"admin": "pwd123",
	})
}
