package main

import "fmt"

type Product struct{
	Name string
	Price float64
	Count int
}

type Service struct{
	Name string
	Price float64
	MinutesWorking int
}

type Maintenance struct{
	Name string
	Price float64
}

func SumProducts(products []Product) float64{
	var total float64
	for _,key:= range products{
		var subtotal = key.Price * float64(key.Count)
		total += subtotal
	}
	return total
}

func SumServices(services []Service) float64{
	var total float64
	for _,key:= range services{
		var minutes int = int(key.MinutesWorking / 30)
		if minutes < 0{
			minutes = 1
		}
		var subtotal = key.Price * cuenta
		total += subtotal
	}
	return total
}

func SumMaintenance(maintenances []Maintenance)float64{
	var total float64
	for _,key:= range products{
		total += key.Price
	}
	return total
}


func main(){


	totalProducts := go SumProducts(products)
	totalServices := go SumServices(services)
	totalMaintenance := go SumMaintenance(maintenances)

	fmt.Println(totalProducts + totalServices + totalMaintenance)


}