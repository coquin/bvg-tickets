package ticket

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Get(id int) (*Ticket, error) {
	return &Ticket{id, 1, "AB"}, nil
}

func (s *Service) Create(userId int, zone string) (int, error) {
	return 1, nil
}
