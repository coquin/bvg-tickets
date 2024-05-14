package ticket

import (
	"bvg-tickets/clock"
	"time"
)

type activation struct {
	createdAt time.Time
	expiresAt time.Time
}

func Activation(clock clock.Clock, duration time.Duration) activation {
	createdAt := clock.Now()
	expiresAt := createdAt.Add(duration)

	return activation{createdAt, expiresAt}
}

func (a activation) ActiveSince() time.Time {
	return a.createdAt
}

func (a activation) ExpiresAt() time.Time {
	return a.expiresAt
}
