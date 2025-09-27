package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goncadanilo/products-api-go/controller"
)

func main() {
	server := gin.Default()

	ProductController := controller.NewProductController()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}
