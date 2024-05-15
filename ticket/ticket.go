package ticket

import (
	"bvg-tickets/clock"
	"bvg-tickets/zone"
	"time"
)

type Validation interface {
	ValidFrom() time.Time
	ValidTill() time.Time
}

type singleTicketValidation struct {
	when    time.Time
	howLong time.Duration
	where   string
}

func (v singleTicketValidation) ValidFrom() time.Time {
	return v.when
}

func (v singleTicketValidation) ValidTill() time.Time {
	return v.when.Add(v.howLong)
}

type ValidatedTicket interface {
	ValidFrom() time.Time
	ValidTill() time.Time
}

type Validator interface {
	Validate() ValidatedTicket
}

type ticket struct {
	Zone zone.Zone
}

type singleTripTicket struct {
	Zone     zone.Zone
	validFor time.Duration
}

type validatedSingleTripTicket struct {
	singleTripTicket
	Validation
	StartLocation string
}

func (t singleTripTicket) Validate(clock clock.Clock, location string) ValidatedTicket {
	validFrom := clock.Now()
	validation := singleTicketValidation{validFrom, t.validFor, location}
	return validatedSingleTripTicket{t, validation, location}
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
