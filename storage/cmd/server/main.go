package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/storage/cmd/server/handler"
	cnn "github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/storage/db"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/storage/internal/products"
)

func main() {
	loadEnv()

	db := cnn.MySQLConnection()
	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	p := handler.NewProduct(serv)

	r := gin.Default()
	pr := r.Group("/api/v1/products")

	pr.POST("/", p.Store())
	pr.GET("/:id", p.GetAll())
	pr.GET("/", p.GetByName())
	pr.PATCH("/:id", p.Update())

	r.Run()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
