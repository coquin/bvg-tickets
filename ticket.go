package bvg_tickets

import "time"

type validation struct {
	validatedAt time.Time
}

type Ticket struct {
	v *validation
}

func New() Ticket {
	return Ticket{nil}
}

func (t *Ticket) IsValid() bool {
	return t.v != nil
}

func (t *Ticket) Validate() {
	t.v = &validation{time.Now()}
}
