package main

import ("fmt"
"errors"
"os")

var err = errors.New("error,el salario ingresado no alcanza el minimo imponible")

func main(){

	salary := 200000
	
	if salary < 150000{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
	
}