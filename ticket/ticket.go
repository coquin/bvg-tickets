package ticket

import "strconv"

type Ticket struct {
	Id     TicketId
	UserId UserId
	Zone   Zone
}

type UserId struct {
	Value int
}

func (i UserId) String() string {
	return strconv.Itoa(i.Value)
}

type TicketId struct {
	Value int
}

func (i TicketId) String() string {
	return strconv.Itoa(i.Value)
}

// Thanks to https://www.reddit.com/r/golang/comments/uvpygm/comment/iab61oh/
type Zone int

const (
	ZoneABC Zone = iota
	ZoneAB
	ZoneBC
)

func (z Zone) String() string {
	return []string{"ABC", "AB", "BC"}[z]
}

type Reader interface {
	Read(TicketId) (*Ticket, error)
}

type Writer interface {
	Write(*Ticket) error
}

type IdGenerator interface {
	// TODO: should be more generic, to allow generating any ids
	NextId() TicketId
}

type Repository interface {
	Reader
	Writer
	IdGenerator
}

type UseCase interface {
	Get(TicketId) (*Ticket, error)
	Purchase(UserId, Zone) (TicketId, error)
}
