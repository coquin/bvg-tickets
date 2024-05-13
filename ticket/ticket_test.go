package ticket_test

import (
	"bvg-tickets/clock"
	"bvg-tickets/ticket"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSingleTicket(t *testing.T) {
	now := time.Now()
	mockClock := clock.Mock(now)
	ticket := ticket.Single(mockClock)

	assert.NotNil(t, ticket)
	assert.Equal(t, now, ticket.CreatedAt())
	assert.False(t, ticket.IsValidated())

	later := now.Add(time.Minute)
	mockClock = clock.Mock(later)
	ticket.Validate(mockClock, "Berlin Hauptbahnhof")

	assert.True(t, ticket.IsValidated())
	assert.Equal(t, later, ticket.ValidFrom())
	assert.Equal(t, "Berlin Hauptbahnhof", ticket.StartLocation())
}

func TestSingleTicketCannotBeValidatedTwice(t *testing.T) {
	now := time.Now()
	mockClock := clock.Mock(now)
	ticket := ticket.Single(mockClock)
	ticket.Validate(mockClock, "Berlin Hauptbahnhof")

	later := now.Add(time.Hour)
	mockClock = clock.Mock(later)
	ticket.Validate(mockClock, "Berlin Ostbahnhof")

	assert.True(t, ticket.IsValidated())
	assert.Equal(t, now, ticket.ValidFrom())
	assert.Equal(t, "Berlin Hauptbahnhof", ticket.StartLocation())
}
