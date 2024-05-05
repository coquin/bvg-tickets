package bvg_tickets_test

import (
	bvgtickets "bvg-tickets"
	"bvg-tickets/clock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTicketNewIsNotValidated(t *testing.T) {
	assert := assert.New(t)
	policy := bvgtickets.Policy{ValidFor: time.Minute}
	ticket := bvgtickets.New(policy)
	mockClock := clock.Mock(time.Now())

	assert.NotNil(ticket)
	assert.False(ticket.IsValid(mockClock))
}

func TestTicketIsValid(t *testing.T) {
	assert := assert.New(t)
	now := time.Now()
	mockClock := clock.Mock(now)
	policy := bvgtickets.Policy{ValidFor: time.Minute * 120}
	ticket := bvgtickets.New(policy)
	ticket.Validate(mockClock)
	assert.True(ticket.IsValid(mockClock))

	now = now.Add(policy.ValidFor + 1)
	mockClock = clock.Mock(now)
	assert.False(ticket.IsValid(mockClock))
}
