package bvg_tickets_test

import (
	bvg_tickets "bvg-tickets"
	"bvg-tickets/clock"
	"bvg-tickets/zone"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTicket(t *testing.T) {
	now := time.Now()
	_clock := clock.Mock(now)
	ticket := bvg_tickets.New(_clock, zone.AB, "Berlin Hauptbahnhof")
	assert.NotNil(t, ticket)
	assert.Equal(t, now, ticket.CreatedAt())
}

func TestSingleTicket(t *testing.T) {
	now := time.Now()
	_clock := clock.Mock(now)
	twoHours := time.Hour * 2
	ticket := bvg_tickets.Single(_clock, zone.AB, twoHours)
	ticket.Validate(_clock, "Berlin Hauptbahnhof")

	assert.Equal(t, now, ticket.CreatedAt())
	assert.Equal(t, now, ticket.ValidFrom())

	twoHoursLater := now.Add(twoHours)
	assert.Equal(t, twoHoursLater, ticket.ValidUntil())
}

func Test24hTicket(t *testing.T) {
	now := time.Now()
	_clock := clock.Mock(now)
	ticket := bvg_tickets.T24h(_clock, zone.AB)
	ticket.Validate(_clock, "")

	assert.Equal(t, now, ticket.CreatedAt())
	assert.Equal(t, now, ticket.ValidFrom())
	assert.Equal(t, now.Add(time.Hour*24), ticket.ValidUntil())
}
