package main

import (
	"fmt"

	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/hackaton-go-bases-main/internal/file"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/hackaton-go-bases-main/internal/service"
)

func main() {
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	f := file.File{Path: "/Users/MARIASODIAZ/github/backpack-bcgow6-mariasol-diazreal/hackaton-go-bases-main/tickets.csv"}

	info, _ := f.Read()
	tickets = append(tickets, info...)
	booking := service.NewBookings(tickets)
	defer func() {
		fmt.Println("ejecucion finalizada")
	}()

	newTicket := service.Ticket{
		Id:          2020,
		Names:       "Sol diaz real",
		Email:       "soldr@email.com",
		Destination: "USA",
		Date:        "11:00",
		Price:       1000,
	}
	_, err := booking.Create(newTicket)
	if err != nil {
		panic(err)
	}
	fmt.Println("Se ha creado tu ticket")
	err = f.Write(newTicket)

	ticket, err := booking.Read(2020)
	if err != nil {
		panic(err)
	}
	fmt.Println(ticket)

	ticket.Price = 666666
	_, err = booking.Update(ticket.Id, ticket)
	if err != nil {
		panic(err)
	}

	fmt.Println("Actualizado: ", ticket)

	_, err = booking.Delete(1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Se ha borrado el ticket")

}
