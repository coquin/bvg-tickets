package ticket

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) Get(id int) (*Ticket, error) {
	return s.repo.Read(id)
}

func (s *Service) Purchase(userId int, zone Zone) (int, error) {
	id := s.repo.NextId()
	t := &Ticket{id, userId, zone}

	if err := s.repo.Write(t); err != nil {
		return 0, err
	}
	return id, nil
}
