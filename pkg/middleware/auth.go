package middleware

import (
	"challenge-api/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc  {
	return func(c *gin.Context) {
		verifyToken, err := helper.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"messager": err.Error(),
			})
			return
		}
		c.Set("UserData", verifyToken)
		c.Next()
	}
}