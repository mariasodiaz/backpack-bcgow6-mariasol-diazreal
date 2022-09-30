package file

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/hackaton-go-bases-main/internal/service"
)

type File struct {
	path string
}

func (f *File) Read() ([]service.Ticket, error) {
	defer func() {
		fmt.Println("ejecucion finalizada")
	}()
	content, err := os.ReadFile("../../tickets.csv")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	var tickets []service.Ticket

	for i := 0; i < len(lines); i++ {
		ticket := strings.Split(lines[i], ",")
		if len(ticket) < 6 {
			err = errors.New("Cantidad insuficiente de parametros")
			return nil, err
		}
		id, _ := strconv.Atoi(ticket[0])
		price, _ := strconv.Atoi(ticket[5])
		tickets[i].Id = id
		tickets[i].Names = ticket[1]
		tickets[i].Email = ticket[2]
		tickets[i].Destination = ticket[3]
		tickets[i].Date = ticket[4]
		tickets[i].Price = price
	}
	return tickets, nil
}

func (f *File) Write(ticket service.Ticket) error {

	ticketString := fmt.Sprintf("%v,%v,%v,%v,%v,%v", ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price)
	ticketsbyte := []byte(ticketString)
	err := os.WriteFile("../../tickets.csv", ticketsbyte, 0644)
	if err != nil {
		return err
	}
	return nil
}
