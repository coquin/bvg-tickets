package bvg_tickets_test

import (
	bvgtickets "bvg-tickets"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTicketNew(t *testing.T) {
	assert := assert.New(t)
	ticket := bvgtickets.New()

	assert.NotNil(ticket)
	assert.False(ticket.IsValid())
}

func TestTicketValidate(t *testing.T) {
	assert := assert.New(t)
	ticket := bvgtickets.New()
	ticket.Validate()

	assert.True(ticket.IsValid())
}
