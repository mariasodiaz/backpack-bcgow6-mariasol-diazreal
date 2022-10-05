package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/goweb/arquitectura/cmd/server/handler"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/goweb/arquitectura/internal/products"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/goweb/arquitectura/pkg/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db := store.New(store.FileType, "./products.json")
	repository := products.NewRepository(db)
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
