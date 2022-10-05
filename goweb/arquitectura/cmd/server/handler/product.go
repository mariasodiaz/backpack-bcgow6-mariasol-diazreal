package handler

import (
	"net/http"
	"os"
	"strconv"

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

var errorPeticion string = "no tienes permisos para la peticion solicitada"

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenReq := context.GetHeader("token")
		if tokenReq != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized, gin.H{"error": errorPeticion})
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
		if tokenReq != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized, gin.H{"error": errorPeticion})
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

func (p *Product) Update() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenReq := context.GetHeader("token")
		if tokenReq != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized, gin.H{"error": errorPeticion})
			return
		}
		id, _ := strconv.Atoi(context.Param("id"))
		var req ProductRequest
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if req.Name == "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "El nombre del producto es requerido"})
			return
		}
		if req.Color == "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "El color del producto es requerido"})
			return
		}
		if req.Price == 0 {
			context.JSON(http.StatusBadRequest, gin.H{"error": "El precio es requerido"})
			return
		}
		if req.Stock == 0 {
			context.JSON(http.StatusBadRequest, gin.H{"error": "El stock es requerido"})
			return
		}
		if req.Code == "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "El codigo es requerido"})
			return
		}
		if req.Date == "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "La fecha es requerida"})
			return
		}
		product, err := p.service.Update(id, req.Name, req.Color, req.Price, req.Stock, req.Code, false, req.Date)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		context.JSON(http.StatusOK, product)
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenReq := context.GetHeader("token")
		if tokenReq != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized, gin.H{"error": errorPeticion})
			return
		}
		id, _ := strconv.Atoi(context.Param("id"))
		err := p.service.Delete(id)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		context.JSON(http.StatusNoContent, gin.H{"data": "el producto se ha eliminado correctamente"})
	}
}

func (p *Product) UpdateMany() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenReq := context.GetHeader("token")
		if tokenReq != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized, gin.H{"error": errorPeticion})
			return
		}
		id, _ := strconv.Atoi(context.Param("id"))
		var req ProductRequest
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product, err := p.service.UpdateMany(id, req.Name, req.Price)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		context.JSON(http.StatusOK, product)
	}
}
