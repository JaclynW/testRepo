package handlers

import (
	"net/http"
	"product-service/models"

	"github.com/gin-gonic/gin"
)

var products = []models.Product{
	{ID: 1, Name: "Laptop", Price: 1000.50},
	{ID: 2, Name: "Mouse", Price: 25.75},
}

func GetAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, product := range products {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}

func AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	products = append(products, product)
	c.JSON(http.StatusOK, product)
}
