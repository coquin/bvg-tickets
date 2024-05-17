package mock

import (
	"bvg-tickets/ticket"
	"fmt"
)

type Repository struct {
	tickets []ticket.Ticket
	nextId  int
}

func NewRepository() *Repository {
	return &Repository{make([]ticket.Ticket, 0), 1}
}

func (r *Repository) Read(id int) (*ticket.Ticket, error) {
	for _, t := range r.tickets {
		if t.Id == id {
			return &t, nil
		}
	}
	return nil, fmt.Errorf("ticket %d not found", id)
}

func (r *Repository) Write(t *ticket.Ticket) error {
	r.tickets = append(r.tickets, *t)
	return nil
}

func (r *Repository) NextId() int {
	nextId := r.nextId
	r.nextId += 1
	return nextId
}
