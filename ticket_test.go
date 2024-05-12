package bvg_tickets_test

import (
	bvg_tickets "bvg-tickets"
	"bvg-tickets/clock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewTicket(t *testing.T) {
	now := time.Now()
	mockClock := clock.Mock(now)
	ticket := bvg_tickets.New(mockClock)
	assert.NotNil(t, ticket)
	assert.Equal(t, now, ticket.CreatedAt())
}
