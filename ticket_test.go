package bvg_tickets_test

import (
	bvgtickets "bvg-tickets"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTicketNew(t *testing.T) {
	assert := assert.New(t)
	ticket := bvgtickets.New()

	assert.NotNil(ticket)
}
