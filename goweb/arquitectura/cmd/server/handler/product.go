package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/goweb/arquitectura/internal/products"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/goweb/arquitectura/pkg/web"
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

// ListProducts godoc
// @Summary     List products
// @Tags        Products
// @Description get products
// @Produce     json
// @Param       token header   string true "token"
// @Success     200   {object} web.Response
// @Failure     401   {object} web.Response
// @Router      /products [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenReq := context.GetHeader("token")
		if tokenReq != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, errorPeticion))
			return
		}
		products, err := p.service.GetAll()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, products, ""))
	}
}
func (p *Product) validar(context *gin.Context, req ProductRequest) string {
	if req.Name == "" {
		return "el nombre es requerido"
	}
	if req.Color == "" {
		return "el color es requerido"
	}
	if req.Price == 0 {
		return "el precio es requerido"
	}
	if req.Stock == 0 {
		return "el stock es requerido"
	}
	if req.Code == "" {
		return "el codigo es requerido"
	}
	if req.Date == "" {
		return "la fecha es requerida"
	}
	return ""
}

// StoreProducts godoc
// @Summary     Store products
// @Tags        Products
// @Description store products
// @Accept      json
// @Produce     json
// @Param       token   header   string  true "token"
// @Param       product body     ProductRequest true "Product to store"
// @Success     200     {object} web.Response
// @Failure     401     {object} web.Response
// @Failure     400     {object} web.Response
// @Router      /products [post]
func (p *Product) Store() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenReq := context.GetHeader("token")
		if tokenReq != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, errorPeticion))
			return
		}
		var req ProductRequest
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		errValidacion := p.validar(context, req)
		if errValidacion != "" {
			context.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, errValidacion))
			return
		}
		product, err := p.service.Store(req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.Date)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, product, ""))
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenReq := context.GetHeader("token")
		if tokenReq != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, errorPeticion))
			return
		}
		id, _ := strconv.Atoi(context.Param("id"))
		var req ProductRequest
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		errValidacion := p.validar(context, req)
		if errValidacion != "" {
			context.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, errValidacion))
			return
		}
		product, err := p.service.Update(id, req.Name, req.Color, req.Price, req.Stock, req.Code, false, req.Date)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, product, ""))
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenReq := context.GetHeader("token")
		if tokenReq != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, errorPeticion))
			return
		}
		id, _ := strconv.Atoi(context.Param("id"))
		err := p.service.Delete(id)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		context.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, ""))
	}
}

func (p *Product) UpdateMany() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenReq := context.GetHeader("token")
		if tokenReq != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, errorPeticion))
			return
		}
		id, _ := strconv.Atoi(context.Param("id"))
		var req ProductRequest
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		product, err := p.service.UpdateMany(id, req.Name, req.Price)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, product, ""))
	}
}
