package incrementing

import (
	counters "github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal"
)

type DefaultService interface {
	Increment(ID string) error
}

type Service struct {
	counters counters.CounterRepository
}

func NewService(cR counters.CounterRepository) Service {
	return Service{counters: cR}
}

func (s Service) Increment(ID string) error {
	counter, err := s.counters.Get(ID)
	if err != nil {
		return err
	}

	counter.Increment()

	err = s.counters.Save(counter)
	if err != nil {
		return err
	}

	return nil
}
