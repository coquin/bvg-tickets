package bvg_tickets

import (
	"bvg-tickets/utils"
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

func (t *Ticket) Validate(clock utils.Clock) {
	t.v = &validation{clock.Now(), clock.Now().Add(t.policy.ValidFor)}
}

func (t *Ticket) IsValid(clock utils.Clock) bool {
	return t.v != nil && !clock.Now().After(t.v.validUntil)
}
