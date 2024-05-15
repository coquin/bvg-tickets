package validation_test

import (
	"bvg-tickets/validation"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidFor(t *testing.T) {
	v := validation.ValidFor{time.Hour * 2}
	assert.NotNil(t, v)
	assert.Equal(t, time.Hour*2, v.Duration)
}

func TestValidForStartingAt(t *testing.T) {
	v := validation.ValidFor{time.Hour * 2}
	lv := validation.ValidForStartingAt{v, "Berlin Hauptbahnhof"}
	assert.NotNil(t, lv)
	assert.Equal(t, time.Hour*2, lv.Duration)
	assert.Equal(t, "Berlin Hauptbahnhof", lv.Location)
}
