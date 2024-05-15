package ticket

import (
	"bvg-tickets/clock"
	"bvg-tickets/zone"
	"time"
)

type ticket struct {
	Zone zone.Zone
}

func New(zone zone.Zone) ticket {
	return ticket{zone}
}

func (t ticket) Validate(clock clock.Clock) validatedTicket {
	return validatedTicket{t, clock.Now()}
}

type validatedTicket struct {
	ticket
	ValidFrom time.Time
}

func (t validatedTicket) AtLocation(location string) validatedTicketAtLocation {
	return validatedTicketAtLocation{t, location}
}

type validatedTicketAtLocation struct {
	validatedTicket
	StartLocation string
}

func SingleTrip(zone zone.Zone, clock clock.Clock, location string) validatedTicketAtLocation {
	return New(zone).Validate(clock).AtLocation(location)
}
