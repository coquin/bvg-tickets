package bvg_tickets_test

import (
	bvg_tickets "bvg-tickets"
	"bvg-tickets/clock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTicket(t *testing.T) {
	now := time.Now()
	mockClock := clock.Mock(now)
	ticket := bvg_tickets.New(mockClock)
	assert.NotNil(t, ticket)
	assert.Equal(t, now, ticket.CreatedAt())
}

func TestNewTimeLimitedTicket(t *testing.T) {
	now := time.Now()
	mockClock := clock.Mock(now)
	twoHours := time.Hour * 2
	ticket := bvg_tickets.NewTimeLimited(mockClock, twoHours)
	ticket.Validate(mockClock)

	assert.Equal(t, now, ticket.CreatedAt())
	assert.Equal(t, now, ticket.ValidFrom())

	twoHoursLater := now.Add(twoHours)
	assert.Equal(t, twoHoursLater, ticket.ValidUntil())
}
