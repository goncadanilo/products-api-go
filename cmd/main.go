package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goncadanilo/products-api-go/controller"
	"github.com/goncadanilo/products-api-go/database"
	"github.com/goncadanilo/products-api-go/repository"
	"github.com/goncadanilo/products-api-go/usecase"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	databaseConnection, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(databaseConnection)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	server.GET("/products", ProductController.GetProducts)
	server.GET("/products/:id", ProductController.GetProductById)
	server.POST("/products", ProductController.CreateProduct)

	server.Run(":8000")
}
