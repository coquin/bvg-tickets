package bvg_tickets

import (
	"bvg-tickets/clock"
	"time"
)

type ticket struct {
	createdAt   time.Time
	validatedAt time.Time
	validFor    time.Duration
}

func New(clock clock.Clock) ticket {
	return ticket{clock.Now(), time.Time{}, 0}
}

func NewTimeLimited(clock clock.Clock, duration time.Duration) ticket {
	return ticket{clock.Now(), time.Time{}, duration}
}

func (t ticket) CreatedAt() time.Time {
	return t.createdAt
}

func (t *ticket) Validate(clock clock.Clock) {
	t.validatedAt = clock.Now()
}

func (t ticket) ValidFrom() time.Time {
	return t.validatedAt
}

func (t ticket) ValidUntil() time.Time {
	return t.validatedAt.Add(t.validFor)
}
