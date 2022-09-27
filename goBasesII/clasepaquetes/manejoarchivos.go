package main

import ("os"
"fmt")

func main(){

	message := "1;233;1\n2;1000,2\n"
	messageBytes := []byte(message)
	os.WriteFile("./productos.csv",messageBytes,0644)


	products,_ := os.ReadFile("./productos.csv")

	productsString := string(products)

	fmt.Printf(productsString)

}