package routes

import (
	user "challenge-api/pkg/controllers"
	// order "challenge-api/pkg/controllers"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

func StartApp() *gin.Engine  {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", user.UserRegister)
		userRouter.POST("login", user.UserLogin)
	}

	return r
}