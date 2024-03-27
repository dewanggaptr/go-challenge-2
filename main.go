package main

import (
	"challenge-api/pkg/database"
	"challenge-api/pkg/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := database.DBInit()

	if err != nil {
		return 
	}
	
	r := gin.Default()

	routes.OrdersRoute(r, db)
	r.Run()
}
