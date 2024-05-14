package ticket

import (
	"bvg-tickets/clock"
	"time"
)

type activation struct {
	createdAt time.Time
}

func Activation(clock clock.Clock) activation {
	return activation{clock.Now()}
}

func (a activation) ActiveSince() time.Time {
	return a.createdAt
}
