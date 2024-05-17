package ticket_test

import (
	"bvg-tickets/ticket"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicketCreate(t *testing.T) {
	var useCase ticket.UseCase = ticket.NewService()

	// Given user creates a ticket
	userId := 1
	zone := "AB"

	// When user creates a ticket
	_ticketId, err := useCase.Create(userId, zone)

	// Then user has the ticket
	assert.NoError(t, err)
	assert.Equal(t, 1, _ticketId)
}

func TestTicketGet(t *testing.T) {
	var useCase ticket.UseCase = ticket.NewService()

	// Given user created a ticket
	userId := 1
	zone := "AB"
	_ticketId, _ := useCase.Create(userId, zone)

	// When user gets the ticket by id
	_ticket, err := useCase.Get(_ticketId)

	// Then ticket is returned
	assert.NoError(t, err)

	expectedTicket := &ticket.Ticket{_ticketId, userId, zone}
	assert.Equal(t, expectedTicket, _ticket)
}
