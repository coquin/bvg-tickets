package ticket

type Ticket struct {
	Id     int
	UserId int
	Zone   string
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
	Purchase(int, string) (int, error)
}
