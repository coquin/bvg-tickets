package mock

import (
	"bvg-tickets/ticket"
	"fmt"
)

type Repository struct {
	tickets []ticket.Ticket
	nextId  ticket.Id
}

func NewRepository() *Repository {
	return &Repository{make([]ticket.Ticket, 0), ticket.Id{Value: 1}}
}

func (r *Repository) Read(id ticket.Id) (*ticket.Ticket, error) {
	for _, t := range r.tickets {
		if t.Id == id {
			return &t, nil
		}
	}
	return nil, fmt.Errorf("ticket %v not found", id)
}

func (r *Repository) Write(t *ticket.Ticket) error {
	r.tickets = append(r.tickets, *t)
	return nil
}

func (r *Repository) NextId() ticket.Id {
	nextId := r.nextId
	r.nextId.Value += 1
	return nextId
}
