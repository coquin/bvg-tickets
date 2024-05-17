package mock

import (
	"bvg-tickets/ticket"
	"fmt"
)

type Repository struct {
	tickets map[ticket.Id]ticket.Ticket
	nextId  ticket.Id
}

func NewRepository() *Repository {
	return &Repository{make(map[ticket.Id]ticket.Ticket), ticket.Id{Value: 1}}
}

func (r *Repository) Read(id ticket.Id) (*ticket.Ticket, error) {
	if t, ok := r.tickets[id]; ok {
		return &t, nil
	}

	return nil, fmt.Errorf("ticket %v not found", id)
}

func (r *Repository) Write(t *ticket.Ticket) error {
	r.tickets[t.Id] = *t
	return nil
}

func (r *Repository) NextId() ticket.Id {
	nextId := r.nextId
	r.nextId.Value += 1
	return nextId
}
