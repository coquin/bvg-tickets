package ticket

type Ticket struct {
	Id     int
	UserId int
	Zone   Zone
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
	Read(int) (*Ticket, error)
}

type Writer interface {
	Write(*Ticket) error
}

type IdGenerator interface {
	NextId() int
}

type Repository interface {
	Reader
	Writer
	IdGenerator
}

type UseCase interface {
	Get(int) (*Ticket, error)
	Purchase(int, Zone) (int, error)
}
