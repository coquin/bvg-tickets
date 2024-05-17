package ticket

import "strconv"

type Ticket struct {
	Id          Id
	UserId      UserId
	Zone        Zone
	IsValidated bool
}

type UserId struct {
	Value int
}

func (i UserId) String() string {
	return strconv.Itoa(i.Value)
}

type Id struct {
	Value int
}

func (i Id) String() string {
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
	Read(Id) (*Ticket, error)
}

type Writer interface {
	Write(*Ticket) error
}

type IdGenerator interface {
	// TODO: should be more generic, to allow generating any ids
	NextId() Id
}

type Repository interface {
	Reader
	Writer
	IdGenerator
}

type UseCase interface {
	Get(Id) (*Ticket, error)
	Purchase(UserId, Zone) (Id, error)
	Validate(Id) error
}
