package main

import "fmt"

func main(){

	var palabra string = "hola"

	fmt.Println(len(palabra))

	for i := 0; i<len(palabra); i++{
		fmt.Printf("%c\n",palabra[i])
	}
}
