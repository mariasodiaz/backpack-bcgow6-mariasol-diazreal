package main

import "fmt"


type Student struct{
	Name string
	LastName string
	DNI string
	Date string
}


func main(){

	student := Student{
		Name: "Sol",
		LastName: "Diaz Real",
		DNI: "41918908",
		Date: "21/11/11",
	}

	fmt.Printf("%+v",student)
}