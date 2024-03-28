package routes

import (
	"log"
	"net/http"

	order "challenge-api/pkg/controllers"
	// user "challenge-api/pkg/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

)

func OrdersRoute(r *gin.Engine, db *gorm.DB) {
	orderHandler := order.NewOrderController(db)
	orderPath := r.Group("/orders")

	orderPath.POST("/", orderHandler.Create)
	orderPath.GET("/", orderHandler.Index)
	orderPath.GET("/:id", orderHandler.Show)
	orderPath.PUT("/:id", orderHandler.Update)
	orderPath.DELETE("/:id", orderHandler.Delete)

	log.Fatal(http.ListenAndServe(":8080", r))
}
