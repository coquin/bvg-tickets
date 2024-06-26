package ticket

import (
	"errors"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) Get(id Id) (*Ticket, error) {
	return s.repo.Read(id)
}

func (s *Service) Purchase(zone Zone) (Id, error) {
	id := s.repo.NextId()
	t := &Ticket{id, zone, false}

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

	if t.IsValidated {
		return errors.New("ticket is already validated")
	}

	t.IsValidated = true

	return s.repo.Write(t)
}
