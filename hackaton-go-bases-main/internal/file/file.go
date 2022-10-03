package file

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/hackaton-go-bases-main/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {

	content, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	tickets := []service.Ticket{}

	for _, value := range lines {
		ticket := strings.Split(value, ",")

		id, err := strconv.Atoi(ticket[0])
		if err != nil {
			return nil, err
		}
		price, err := strconv.Atoi(ticket[5])
		if err != nil {
			return nil, err
		}
		newTicket := service.Ticket{
			Id:          id,
			Names:       ticket[1],
			Email:       ticket[2],
			Destination: ticket[3],
			Date:        ticket[4],
			Price:       price,
		}
		tickets = append(tickets, newTicket)
	}
	return tickets, nil
}

func (f *File) Write(ticket service.Ticket) error {

	file, err := os.OpenFile(f.Path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	ticketString := fmt.Sprintf("\n%v,%v,%v,%v,%v,%v", ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price)
	ticketsByte := []byte(ticketString)

	_, errWrite := file.Write(ticketsByte)
	if errWrite != nil {
		fmt.Println(err)
	}
	file.Close()
	if err != nil {
		return err
	}
	return nil
}
