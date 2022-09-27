package main

import "fmt"

type Matrix struct{
	Valores [][]int
	Alto int
	Ancho int
	Cuadratica bool
	ValorMaximo int
}

func (m Matrix) Set(numeros...float64){
	for _,key:= range numeros{
		append(m,key)
	}
}

func (m Matrix) Print(){
	for i:=0; i<len(m.Alto);i++{
		for j:=0; j<len(m.Ancho);j++{
			fmt.Printf("%i",m.Valores[i][j])
		}
		fmt.Print("\n")
	}
}

func main(){

	m := Matrix{
		Alto: 10,
		Ancho: 10,
		Cuadratica: true,
		ValorMaximo: 10,
	}
	m.Set(4.4,3.0,5.3)
	m.Print()
}