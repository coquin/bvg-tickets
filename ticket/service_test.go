package ticket_test

import (
	"bvg-tickets/ticket"
	"bvg-tickets/ticket/mock"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicketPurchase(t *testing.T) {
	repo := mock.NewRepository()
	var useCase ticket.UseCase = ticket.NewService(repo)

	// Given user wants to purchase a ticket for AB zone
	userId := ticket.UserId{1}
	zone := ticket.ZoneAB

	// When user purchases a ticket
	_ticketId, err := useCase.Purchase(userId, zone)

	// Then user has the ticket
	assert.NoError(t, err)
	assert.Equal(t, ticket.Id{1}, _ticketId)
}

func TestTicketGet(t *testing.T) {
	repo := mock.NewRepository()
	var useCase ticket.UseCase = ticket.NewService(repo)

	// Given user purchased a ticket
	userId := ticket.UserId{1}
	zone := ticket.ZoneAB
	_ticketId, _ := useCase.Purchase(userId, zone)

	// When user gets the ticket by id
	_ticket, err := useCase.Get(_ticketId)

	// Then ticket is returned
	assert.NoError(t, err)

	expectedTicket := &ticket.Ticket{_ticketId, userId, zone, false}
	assert.Equal(t, expectedTicket, _ticket)
}

func TestTicketGetNotFound(t *testing.T) {
	repo := mock.NewRepository()
	var useCase ticket.UseCase = ticket.NewService(repo)

	// Given ticket does not exist
	id := ticket.Id{1}

	// When getting ticket
	_ticket, err := useCase.Get(id)

	// Then an error is returned
	expectedError := fmt.Errorf("ticket 1 not found")
	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}
	assert.Nil(t, _ticket)
}

func TestTicketValidate(t *testing.T) {
	repo := mock.NewRepository()
	var useCase ticket.UseCase = ticket.NewService(repo)

	// Given user purchased a ticket
	userId := ticket.UserId{1}
	zone := ticket.ZoneAB
	_ticketId, _ := useCase.Purchase(userId, zone)

	// When user validates a ticket
	err := useCase.Validate(_ticketId)

	// Then the ticket is validated
	assert.NoError(t, err)

	_ticket, _ := useCase.Get(_ticketId)
	expectedTicket := &ticket.Ticket{_ticketId, userId, zone, true}
	assert.Equal(t, expectedTicket, _ticket)
}
