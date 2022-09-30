package service

import (
	"errors"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	if (t == Ticket{}) {
		return t, errors.New("ticket is empty")
	}
	return Ticket{
		Id:          t.Id,
		Names:       t.Names,
		Email:       t.Email,
		Destination: t.Destination,
		Date:        t.Date,
		Price:       t.Price,
	}, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	for _, element := range (*b).Tickets {
		if element.Id == id {
			return element, nil
		}
	}
	err := errors.New("no se encontro ningun ticket con ese id")
	return Ticket{}, err
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	_, err := b.Read(id)
	if err != nil {
		return Ticket{}, errors.New("no se encontro el id a actualizar")
	}

	return t, nil
}

func (b *bookings) Delete(id int) (int, error) {
	_, err := b.Read(id)
	if err != nil {
		return 0, errors.New("no se encontro el id a actualizar")
	}
	(*b).Tickets = append((*b).Tickets[:id], (*b).Tickets[id+1:]...)
	return id, nil
}
