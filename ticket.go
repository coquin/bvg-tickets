package bvg_tickets

import (
	"bvg-tickets/clock"
	"time"
)

type validation struct {
	validatedAt time.Time
	validUntil  time.Time
}

type Ticket struct {
	policy Policy
	v      *validation
}

func New(policy Policy) Ticket {
	return Ticket{policy, nil}
}

func (t *Ticket) Validate(clock clock.Clock) {
	now := clock.Now()
	t.v = &validation{now, now.Add(t.policy.ValidFor)}
}

func (t *Ticket) IsValid(clock clock.Clock) bool {
	return t.v != nil && !clock.Now().After(t.v.validUntil)
}
