package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/goweb/arquitectura/cmd/server/handler"

	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/goweb/arquitectura/docs"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/goweb/arquitectura/internal/products"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/goweb/arquitectura/pkg/store"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Bootcamp Go Wave 6 - API
// @version         1.0
// @description     This is a simple API development by Digital House's team.
// @termsOfService  https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name   API Support Digital House
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
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

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(files.Handler))

	r := router.Group("./products")
	r.GET("/", product.GetAll())
	r.POST("/", product.Store())
	r.PUT("/:id", product.Update())
	r.DELETE("/:id", product.Delete())
	r.PATCH("/:id", product.UpdateMany())

	router.Run()
}
