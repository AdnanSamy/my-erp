package main

import (
	"fmt"
	"sam/my-erp/config"
	"sam/my-erp/controller"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("can't load .env file. error: %v", err)
	}

	contactController := controller.GetContactController()

	r := gin.Default()

	// connect to DB
	if err := config.ConnectDB(); err != nil {
		fmt.Printf("failed database setup. error: %v", err)
	}

	v1 := r.Group("/v1")
	{
		v1.GET("/contact", contactController.All)
		v1.POST("/contact", contactController.Create)
	}

	r.Run("localhost:8000")
}
