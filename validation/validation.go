package validation

import (
	"bvg-tickets/clock"
	"time"
)

type Validated interface {
	ValidFrom() time.Time
	ValidTill() time.Time
}

type ValidatedAt interface {
	Validated
	ValidatedAt() string
}

type ValidationRule interface {
	ValidTill(time.Time) time.Time
}

type ValidFor struct {
	Duration time.Duration
}

func (v ValidFor) ValidTill(now time.Time) time.Time {
	return now.Add(v.Duration)
}

func Validate(clock clock.Clock, vr ValidationRule) Validated {
	validFrom := clock.Now()
	validTill := vr.ValidTill(validFrom)

	return validation{validFrom, validTill}
}

func ValidateAt(clock clock.Clock, vr ValidationRule, location string) ValidatedAt {
	validFrom := clock.Now()
	validTill := vr.ValidTill(validFrom)
	v := validation{validFrom, validTill}

	return validationAt{v, location}
}

type validation struct {
	validFrom time.Time
	validTill time.Time
}

func (v validation) ValidFrom() time.Time {
	return v.validFrom
}

func (v validation) ValidTill() time.Time {
	return v.validTill
}

type validationAt struct {
	validation
	startLocation string
}

func (v validationAt) ValidatedAt() string {
	return v.startLocation
}
