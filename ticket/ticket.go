package ticket

import (
	"bvg-tickets/zone"
	"time"
)

type ticket struct {
	zone zone.Zone
}

type multiTripTicket struct {
	ticket
	ValidFor time.Duration
}

type singleTripTicket struct {
	ticket
	ValidFor time.Duration
}

func Ticket(zone zone.Zone) ticket {
	return ticket{zone}
}

func (t ticket) FareZone() zone.Zone {
	return t.zone
}

func (t ticket) MultiTrip(validFor time.Duration) multiTripTicket {
	return multiTripTicket{t, validFor}
}

func (t ticket) SingleTrip(validFor time.Duration) singleTripTicket {
	return singleTripTicket{t, validFor}
}
