package main

import ("fmt" 
"errors")

func promedio(notas... int) (int,error){
	if len(notas) == 0{
		return 0,errors.New("No se puede dividir por 0")
	}
	var suma int
	for _,value:= range notas{
		suma += value
	}
	return suma/len(notas),nil //si no escribo nunca el error, el nil ya esta
}

func main(){

	fmt.Println(promedio(4,5,3,2,5))
}