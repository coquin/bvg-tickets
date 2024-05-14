package ticket_test

import (
	"bvg-tickets/clock"
	"bvg-tickets/ticket"
	"bvg-tickets/zone"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTicket(t *testing.T) {
	_ticket := ticket.New(zone.AB)
	assert.NotNil(t, _ticket)
	assert.Equal(t, zone.AB, _ticket.Zone)

	now := time.Now()
	_clock := clock.Mock(now)
	_validatedTicket := _ticket.Validate(_clock)
	assert.Equal(t, now, _validatedTicket.ValidFrom)

	startLocation := "Berlin Hauptbahnhof"
	_ticketValidatedAtLocation := _validatedTicket.AtLocation(startLocation)
	assert.Equal(t, startLocation, _ticketValidatedAtLocation.StartLocation)
}
