package main

import "github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/hackaton-go-bases-main/internal/service"

func main() {
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)
}
