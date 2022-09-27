package main

import "fmt"


func main(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println(employees["Benjamin"])
	var count int = 0

	for _,value := range employees {
		if( value > 25){
			count++
		}
	}

	fmt.Println(count)

	employees["Federico"] = 25

	fmt.Println(employees)

	delete(employees,"Pedro")

	fmt.Println(employees)
}