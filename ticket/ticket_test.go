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
}
