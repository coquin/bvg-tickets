package bvg_tickets

import (
	"bvg-tickets/clock"
	"bvg-tickets/zone"
	"time"
)

type ticket struct {
	createdAt     time.Time
	validFrom     time.Time
	validFor      time.Duration
	zone          zone.Zone
	startLocation string
}

// Deprecated
// TODO: remove
func New(clock clock.Clock, zone zone.Zone, start string) ticket {
	return ticket{clock.Now(), time.Time{}, 0, zone, start}
}

func Single(clock clock.Clock, zone zone.Zone, duration time.Duration, start string) ticket {
	return ticket{clock.Now(), time.Time{}, duration, zone, start}
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
