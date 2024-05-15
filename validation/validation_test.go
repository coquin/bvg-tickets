package validation_test

import (
	"bvg-tickets/clock"
	"bvg-tickets/validation"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidFor(t *testing.T) {
	v := validation.ValidFor{time.Hour * 2}
	assert.NotNil(t, v)
	assert.Equal(t, time.Hour*2, v.Duration)

	now := time.Now()
	_clock := clock.Mock(now)
	_validation := validation.Validate(_clock, v)

	assert.Equal(t, now, _validation.ValidFrom())
	assert.Equal(t, now.Add(time.Hour*2), _validation.ValidTill())
}

func TestValidForStartingAt(t *testing.T) {
	v := validation.ValidFor{time.Hour * 2}
	now := time.Now()
	_clock := clock.Mock(now)
	_validation := validation.ValidateAt(_clock, v, "Berlin Hauptbahnhof")

	assert.Equal(t, now, _validation.ValidFrom())
	assert.Equal(t, now.Add(time.Hour*2), _validation.ValidTill())
	assert.Equal(t, "Berlin Hauptbahnhof", _validation.ValidatedAt())
}
