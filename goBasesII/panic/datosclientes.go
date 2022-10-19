package main

import ("fmt"
"os")


func main(){
	var archivo string = "customers.txt"

	defer func (){
		fmt.Println("ejecucion finalizada")
	}()

	_,err := os.ReadFile(archivo)
	if(err != nil){
		panic("el archivo indicado no fue encontrado o está dañado")
	}
}