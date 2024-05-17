package ticket

type Ticket struct {
	Id     int
	UserId int
	Zone   string
}

type UseCase interface {
	Get(int) (*Ticket, error)
	Create(int, string) (int, error)
}
