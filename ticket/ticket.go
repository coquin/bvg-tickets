package ticket

import (
	"bvg-tickets/clock"
	"bvg-tickets/zone"
	"time"
)

type ticket struct {
	Zone zone.Zone
}

type singleTripTicket struct {
	Zone     zone.Zone
	validFor time.Duration
}

type validatedSingleTripTicket struct {
	singleTripTicket
	ValidFrom     time.Time
	ValidUntil    time.Time
	StartLocation string
}

func (t singleTripTicket) Validate(clock clock.Clock, location string) validatedSingleTripTicket {
	validFrom := clock.Now()
	validUntil := validFrom.Add(time.Hour * 2)
	return validatedSingleTripTicket{t, validFrom, validUntil, location}
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

func SingleTrip(zone zone.Zone) singleTripTicket {
	return singleTripTicket{zone, time.Hour * 2}
}
