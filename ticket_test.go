package bvg_tickets_test

import (
	bvgtickets "bvg-tickets"
	"bvg-tickets/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTicketNewIsNotValidated(t *testing.T) {
	assert := assert.New(t)
	policy := bvgtickets.Policy{ValidFor: time.Minute}
	ticket := bvgtickets.New(policy)
	clock := utils.NewMockClock(time.Now())

	assert.NotNil(ticket)
	assert.False(ticket.IsValid(clock))
}

func TestTicketIsValid(t *testing.T) {
	assert := assert.New(t)
	now := time.Now()
	clock := utils.NewMockClock(now)
	policy := bvgtickets.Policy{ValidFor: time.Minute * 120}
	ticket := bvgtickets.New(policy)
	ticket.Validate(clock)
	assert.True(ticket.IsValid(clock))

	now = now.Add(policy.ValidFor + 1)
	clock = utils.NewMockClock(now)
	assert.False(ticket.IsValid(clock))
}
