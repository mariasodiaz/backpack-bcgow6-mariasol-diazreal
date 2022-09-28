package main

import ("fmt"
"os")

func main(){

	salary := 200000
	
	if salary < 150000{
		err := fmt.Errorf("error,el salario ingresado no alcanza el minimo imponible")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
	
}