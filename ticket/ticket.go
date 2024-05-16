package ticket

import (
	"bvg-tickets/zone"
	"time"
)

type Ticket struct {
	Zone     zone.Zone
	ValidFor time.Duration
}

type SingleTripTicket struct {
	Ticket
}
