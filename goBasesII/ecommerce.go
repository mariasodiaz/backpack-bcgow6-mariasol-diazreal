package main

import "fmt"

type User struct{
	Name string
	LastName string
	Email string
	Products []Product
}

type Product struct{
	Name string
	Price float64
	Count int
}

func NewProduct(name string, price float64, product *Product){
	(*product).Name = name
	(*product).Price = price
}

func addProduct(user *User,product Product, count int){
	product.Count = count
	(*user).Products = append((*user).Products,product)
}

func DeleteProducts(user *User){
	(*user).Products = nil
}

func main(){
	var user User
	var product Product
	NewProduct("coca",75.0,&product)
	user.Name = "Sol"
	user.LastName = "Diaz"
	user.Email = "soldiazreal"
	addProduct(&user,product,3)
	
	fmt.Println(user)

}