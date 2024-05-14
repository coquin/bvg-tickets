package ticket_test

import (
	"bvg-tickets/clock"
	"bvg-tickets/ticket"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestActivationCreated(t *testing.T) {
	now := time.Now()
	_clock := clock.Mock(now)
	_a := ticket.Activation(_clock, time.Hour*2)

	assert.Equal(t, now, _a.ActiveSince())
	assert.Equal(t, now.Add(time.Hour*2), _a.ExpiresAt())
}
