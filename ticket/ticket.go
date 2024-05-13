package ticket

import (
	"bvg-tickets/clock"
	"bvg-tickets/zone"
	"time"
)

type ticket struct {
	createdAt  time.Time
	validation validation
	zone       zone.Zone
}

type validation struct {
	createdAt time.Time
	location  string
}

const ValidDuration time.Duration = time.Hour * 2

var emptyValidation validation = validation{time.Time{}, ""}

func Single(clock clock.Clock, zone zone.Zone) ticket {
	return ticket{clock.Now(), emptyValidation, zone}
}

func (t ticket) CreatedAt() time.Time {
	return t.createdAt
}

func (t ticket) IsValidated() bool {
	return !t.validation.createdAt.Before(t.createdAt)
}

func (t *ticket) Validate(clock clock.Clock, startLocation string) {
	if t.IsValidated() {
		return
	}

	t.validation = validation{clock.Now(), startLocation}
}

func (t ticket) ValidFrom() time.Time {
	return t.validation.createdAt
}

func (t ticket) ValidUntil() time.Time {
	return t.validation.createdAt.Add(ValidDuration)
}

func (t ticket) StartLocation() string {
	return t.validation.location
}
