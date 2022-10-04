package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	Price     int    `json:"price"`
	Stock     int    `json:"stock"`
	Code      string `json:"code"`
	Published bool   `json:"published"`
	Date      string `json:"date"`
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
		return nil, errors.New("hubo un error al leer el archivo")
	}
	json.Unmarshal(raw, &products)
	return products, nil
}

func GetAll1(context *gin.Context) {
	products, _ := GetProducts()
	var filtrados []*Product

	for _, value := range products {
		price, _ := strconv.Atoi(context.Query("price"))
		stock, _ := strconv.Atoi(context.Query("stock"))
		published, _ := strconv.ParseBool(context.Query("published"))
		if context.Query("name") == value.Name && context.Query("color") == value.Color && price == value.Price && stock == value.Stock && context.Query("code") == value.Code && published == value.Published && context.Query("date") == value.Date {
			filtrados = append(filtrados, &value)
		}
	}
	if len(filtrados) != 0 {
		context.JSON(200, filtrados)
	}
}

func GetById(context *gin.Context) {
	products, _ := GetProducts()
	id, _ := strconv.Atoi(context.Param("id"))

	for _, value := range products {
		if value.Id == id {
			context.JSON(200, value)
			return
		}
	}
	context.JSON(404, "id no encontrado")

}
func main() {
	router := gin.Default()
	router.GET("./saludar", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hola Sol",
		})
	})

	router.GET("./productos", GetAll1)
	router.GET("./productos/:id", GetById)

	router.Run()
}
