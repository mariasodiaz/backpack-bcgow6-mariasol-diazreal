package main

import "fmt"

func main(){
	var edad int = 25
	var esEmpleado bool = true
	var antiguedad int = 5
	var sueldo int = 1000

	switch{
	case edad > 22 && esEmpleado && antiguedad > 1:
		if sueldo > 100000{
			fmt.Println("Te otorgaremos prestamos sin intereses")
		} else {
			fmt.Println("Te otorgaremos prestamos con intereses")
		}
	default:
		fmt.Println("No podemos otorgarte ningun prestamo")
	}

}