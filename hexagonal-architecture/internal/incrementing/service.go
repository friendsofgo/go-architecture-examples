package incrementing

import (
	counters "github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal"
)

type Service interface {
	Increment(ID string) error
}

type DefaultService struct {
	counters counters.CounterRepository
}

func NewService(cR counters.CounterRepository) DefaultService {
	return DefaultService{counters: cR}
}

func (s DefaultService) Increment(ID string) error {
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
