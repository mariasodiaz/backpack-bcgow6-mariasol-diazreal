package main

import "fmt"


type product struct{
	tipo string
	nombre string
	precio float64
}

type tienda struct{
	products []product
}

type Product interface{
	CalcularCosto()
}

type Ecommerce interface{
	Total()
	Agregar()
}

func newProduct(tipo,nombre string,precio float64)Product{
	p := product
	p.tipo = tipo
	p.nombre = nombre
	p.precio = precio

	return &p{tipo,nombre,precio}
}

func newTienda(products ...string)Ecommerce{
	&tienda{products}
}

func main(){



}