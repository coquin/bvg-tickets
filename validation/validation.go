package validation

import "time"

// type Interface interface {

// }

type ValidFor struct {
	Duration time.Duration
}

type ValidForStartingAt struct {
	ValidFor
	Location string
}
