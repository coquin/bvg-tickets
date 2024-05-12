package bvg_tickets

import (
	"bvg-tickets/clock"
	"time"
)

type ticket struct {
	createdAt time.Time
}

func New(clock clock.Clock) ticket {
	return ticket{clock.Now()}
}

func (t ticket) CreatedAt() time.Time {
	return t.createdAt
}
