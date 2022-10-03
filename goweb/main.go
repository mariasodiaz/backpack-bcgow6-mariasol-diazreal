package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id        int
	Name      string
	Color     string
	Price     int
	Stock     int
	Code      string
	Published bool
	Date      string
}

var products = []Product{
	{Id: 1, Name: "Cama", Color: "Blanco", Price: 150000, Stock: 5, Code: "AF289A", Published: true, Date: "20/09/2022"},
	{Id: 2, Name: "Televisor", Color: "Negro", Price: 200000, Stock: 1, Code: "AF289A", Published: false, Date: "22/09/2022"},
	{Id: 3, Name: "Cocina", Color: "Plateado", Price: 80000, Stock: 3, Code: "BE224A", Published: true, Date: "01/01/2022"},
	{Id: 4, Name: "Plancha", Color: "Blanco", Price: 200000, Stock: 10, Code: "CL208D", Published: true, Date: "15/04/2022"},
}

func GetAll(context *gin.Context) {
	context.JSON(200, products)
}

func GetProducts() ([]Product, error) {
	var products []Product
	raw, err := ioutil.ReadFile("./products.json")
	if err != nil {
		return nil, errors.New("Hubo un error al leer el archivo")
	}
	json.Unmarshal(raw, &products)
	return products, nil
}

func GetAll1(context *gin.Context) {
	products, err := GetProducts()
	if err == nil {
		context.JSON(200, products)
	}
}

func main() {
	router := gin.Default()
	router.GET("./saludar", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hola Sol",
		})
	})

	router.GET("./productos", GetAll1)

	router.Run()
}
