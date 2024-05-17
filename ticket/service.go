package ticket

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) Get(id Id) (*Ticket, error) {
	return s.repo.Read(id)
}

func (s *Service) Purchase(userId UserId, zone Zone) (Id, error) {
	id := s.repo.NextId()
	t := &Ticket{id, userId, zone, false}

	if err := s.repo.Write(t); err != nil {
		return Id{0}, err
	}
	return id, nil
}

func (s *Service) Validate(id Id) error {
	t, err := s.repo.Read(id)

	if err != nil {
		return err
	}

	t.IsValidated = true

	return s.repo.Write(t)
}
