package ticket

import (
	"bvg-tickets/clock"
	"time"
)

type ticket struct {
	createdAt     time.Time
	validatedAt   time.Time
	startLocation string
}

func Single(clock clock.Clock) ticket {
	return ticket{clock.Now(), time.Time{}, ""}
}

func (t ticket) CreatedAt() time.Time {
	return t.createdAt
}

func (t ticket) IsValidated() bool {
	return !t.validatedAt.Before(t.createdAt)
}

func (t *ticket) Validate(clock clock.Clock, startLocation string) {
	t.validatedAt = clock.Now()
	t.startLocation = startLocation
}

func (t ticket) ValidFrom() time.Time {
	return t.validatedAt
}

func (t ticket) StartLocation() string {
	return t.startLocation
}
