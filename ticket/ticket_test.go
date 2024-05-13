package ticket_test

import (
	"bvg-tickets/clock"
	"bvg-tickets/ticket"
	"bvg-tickets/zone"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSingleTicket(t *testing.T) {
	now := time.Now()
	mockClock := clock.Mock(now)
	_ticket := ticket.Single(mockClock, zone.AB)

	assert.NotNil(t, _ticket)
	assert.Equal(t, now, _ticket.CreatedAt())
	assert.False(t, _ticket.IsValidated())

	later := now.Add(time.Minute)
	mockClock = clock.Mock(later)
	err := _ticket.Validate(mockClock, "Berlin Hauptbahnhof")

	if assert.NoError(t, err) {
		assert.True(t, _ticket.IsValidated())
		assert.Equal(t, later, _ticket.ValidFrom())
		assert.Equal(t, later.Add(ticket.ValidDuration), _ticket.ValidUntil())
		assert.Equal(t, "Berlin Hauptbahnhof", _ticket.StartLocation())
	}
}

func TestSingleTicketCannotBeValidatedTwice(t *testing.T) {
	now := time.Now()
	mockClock := clock.Mock(now)
	_ticket := ticket.Single(mockClock, zone.AB)
	_ticket.Validate(mockClock, "Berlin Hauptbahnhof")

	later := now.Add(time.Hour)
	mockClock = clock.Mock(later)

	err := _ticket.Validate(mockClock, "Berlin Ostbahnhof")
	assert.Error(t, err)

	assert.True(t, _ticket.IsValidated())
	assert.Equal(t, now, _ticket.ValidFrom())
	assert.Equal(t, now.Add(ticket.ValidDuration), _ticket.ValidUntil())
	assert.Equal(t, "Berlin Hauptbahnhof", _ticket.StartLocation())
}
