package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/goweb/arquitectura/cmd/server/handler"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/goweb/arquitectura/internal/products"
)

func main() {
	repository := products.NewRepository()
	service := products.NewService(repository)

	product := handler.NewProduct(service)

	router := gin.Default()

	r := router.Group("./products")
	r.GET("/", product.GetAll())
	r.POST("/", product.Store())
	r.PUT("/:id", product.Update())
	r.DELETE("/:id", product.Delete())
	r.PATCH("/:id", product.UpdateMany())

	router.Run()
}
