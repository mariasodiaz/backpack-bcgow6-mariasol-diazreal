package server

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

	router.GET("./products", product.GetAll())
	router.POST("./products", product.Store())
}
