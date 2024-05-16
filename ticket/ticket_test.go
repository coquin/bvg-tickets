package ticket_test

import (
	"bvg-tickets/ticket"
	"bvg-tickets/zone"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTicket(t *testing.T) {
	_24h := time.Hour * 24
	_ticket := ticket.Ticket{zone.AB, _24h}

	assert.Equal(t, zone.AB, _ticket.Zone)
	assert.Equal(t, _24h, _ticket.ValidFor)
}

func TestSingleTripTicket(t *testing.T) {
	_ticket := ticket.SingleTripTicket{ticket.Ticket{zone.AB, time.Hour * 2}}

	assert.Equal(t, zone.AB, _ticket.Zone)
	assert.Equal(t, time.Hour*2, _ticket.ValidFor)
}
