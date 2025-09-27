package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goncadanilo/products-api-go/model"
)

type productController struct {
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:          1,
			Name:        "Teclado mecânico",
			Description: "Teclado mecânico Keychron k3 v2",
			Price:       800,
			Slug:        "teclado-mecanico",
		},
	}

	ctx.JSON(http.StatusOK, products)
}
