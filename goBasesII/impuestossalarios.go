package main

import "fmt"

func impuestoSalario(salario int)int{
	if salario > 150000{
		return 27
	}else if salario > 50000{
		return 17
	}else{
		return 0
	}
}

func main(){

	impuesto := impuestoSalario(150040)

	fmt.Printf("El impuesto es de: %v",impuesto)
}