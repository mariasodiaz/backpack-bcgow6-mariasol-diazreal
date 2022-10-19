package main

import ("fmt"
"math/rand"
"errors"
"time"
"os")

type Customer struct{
	Legajo int
	Nombre string
	Dni string
	Telefono string
	Domicilio string
}

func generarLegajo()int{
	return 0 + rand.Intn(1000000-0)
}

func validar(customer Customer)(ok string, err error) {
	if customer.Legajo == 0{
		return "",errors.New("Legajo esta vacio")
	}
	if customer.Nombre == ""{
		return "",errors.New("Nombre esta vacio")
	}
	if customer.Dni == ""{
		return "",errors.New("documento esta vacio")
	}
	if customer.Telefono == ""{
		return "",errors.New("telefono esta vacio")
	}
	if customer.Domicilio == ""{
		return "",errors.New("domicilio esta vacio")
	}
	return "ok",nil
}

func main(){
	rand.Seed(time.Now().UnixNano())
	legajo := generarLegajo()
	var customer Customer
	customer.Legajo = legajo

	var archivo string = "customers.txt"

	defer func (){
		fmt.Println("ejecucion finalizada")
		err := recover()
		if err != nil{
			fmt.Println("se detectaron varios errores en tiempo de ejecucion")
		}
		fmt.Println("No han quedado archivos abiertos")
	}()

	_,err := os.ReadFile(archivo)
	if(err != nil){
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	_,err = validar(customer)
	if err != nil{
		panic(err)
	}
}