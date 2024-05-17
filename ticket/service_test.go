package ticket_test

import (
	"bvg-tickets/ticket"
	"bvg-tickets/ticket/mock"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicketCreate(t *testing.T) {
	repo := mock.NewRepository()
	var useCase ticket.UseCase = ticket.NewService(repo)

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
	repo := mock.NewRepository()
	var useCase ticket.UseCase = ticket.NewService(repo)

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

func TestTicketGetNotFound(t *testing.T) {
	repo := mock.NewRepository()
	var useCase ticket.UseCase = ticket.NewService(repo)

	// Given ticket does not exist
	id := 1

	// When getting ticket
	_ticket, err := useCase.Get(id)

	// Then an error is returned
	expectedError := fmt.Errorf("ticket 1 not found")
	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}
	assert.Nil(t, _ticket)
}
