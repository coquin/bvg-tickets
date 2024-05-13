package bvg_tickets

import (
	"bvg-tickets/clock"
	"time"
)

type ticket struct {
	createdAt time.Time
	validFrom time.Time
	validFor  time.Duration
	zone      Zone
}

func New(clock clock.Clock, zone Zone) ticket {
	return ticket{clock.Now(), time.Time{}, 0, zone}
}

func NewTimeLimited(clock clock.Clock, zone Zone, duration time.Duration) ticket {
	return ticket{clock.Now(), time.Time{}, duration, zone}
}

func (t ticket) CreatedAt() time.Time {
	return t.createdAt
}

func (t *ticket) Validate(clock clock.Clock) {
	t.validFrom = clock.Now()
}

func (t ticket) ValidFrom() time.Time {
	return t.validFrom
}

func (t ticket) ValidUntil() time.Time {
	return t.validFrom.Add(t.validFor)
}
