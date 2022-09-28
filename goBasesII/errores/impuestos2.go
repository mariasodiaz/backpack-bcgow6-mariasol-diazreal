package main

import ("fmt"
"errors"
"os")

func main(){

	salary := 200000
	
	if salary < 150000{
		fmt.Println(errors.New("error,el salario ingresado no alcanza el minimo imponible"))
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
	
}