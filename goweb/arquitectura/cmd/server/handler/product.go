package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/goweb/arquitectura/internal/products"
)

type ProductRequest struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	Price     int    `json:"price"`
	Stock     int    `json:"stock"`
	Code      string `json:"code"`
	Published bool   `json:"published"`
	Date      string `json:"date"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenReq := context.GetHeader("token")
		if tokenReq != "12345" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "no tienes permisos para la peticion solicitada"})
			return
		}
		products, err := p.service.GetAll()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		context.JSON(http.StatusOK, products)
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenReq := context.GetHeader("token")
		if tokenReq != "12345" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "no tienes permisos para la peticion solicitada"})
			return
		}
		var req ProductRequest
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "hay un campo vacio"})
			return
		}
		product, err := p.service.Store(req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.Date)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, product)
	}
}
