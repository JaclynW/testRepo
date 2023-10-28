package main

import (
	"product-service/services/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/products", handlers.GetAllProducts)
	r.GET("/products/:id", handlers.GetProductByID)
	r.POST("/products", handlers.AddProduct)

	r.Run(":8080")
}
