package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PedidoRequest struct {
	Id        int    `json:"id"`
	Name      string `json:"name" binding:"required"`
	Color     string `json:"color" binding:"required"`
	Price     int    `json:"price" binding:"required"`
	Stock     int    `json:"stock" binding:"required"`
	Code      string `json:"code" binding:"required"`
	Published bool   `json:"published" binding:"required"`
	Date      string `json:"date" binding:"required"`
}

var requests []PedidoRequest

var token string = "123456"

func PostProduct(context *gin.Context) {
	var req PedidoRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "hay un campo vacio"})
		return
	}

	tokenReq := context.GetHeader("token")
	if tokenReq != token {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "no tienes permisos para la peticion solicitada"})
		return
	}
	var id int = 0
	if len(requests) > 0 {
		id = requests[len(requests)-1].Id
	}
	req.Id = id + 1
	requests = append(requests, req)
	context.JSON(http.StatusOK, req)

}

func main() {
	router := gin.Default()

	router.POST("./productos", PostProduct)

	router.Run()
}
