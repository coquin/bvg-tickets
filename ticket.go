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

const (
	d24h time.Duration = time.Hour * 24
)

// Deprecated
// TODO: remove
func New(clock clock.Clock, zone zone.Zone, start string) ticket {
	return ticket{clock.Now(), time.Time{}, 0, zone, start}
}

func Single(clock clock.Clock, zone zone.Zone, duration time.Duration) ticket {
	return ticket{clock.Now(), time.Time{}, duration, zone, ""}
}

func T24h(clock clock.Clock, zone zone.Zone) ticket {
	return ticket{clock.Now(), time.Time{}, d24h, zone, ""}
}

func (t ticket) CreatedAt() time.Time {
	return t.createdAt
}

func (t *ticket) Validate(clock clock.Clock, startLocation string) {
	t.validFrom = clock.Now()
	t.startLocation = startLocation
}

func (t ticket) ValidFrom() time.Time {
	return t.validFrom
}

func (t ticket) ValidUntil() time.Time {
	return t.validFrom.Add(t.validFor)
}
