package incrementer

import (
	counters "github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/errors"
)

type Service interface {
	Increment(ID string) error
}

type service struct {
	counters  counters.CounterRepository
}

func NewService(cR counters.CounterRepository) Service {
	return &service{counters: cR}
}

func (s *service) Increment(ID string) error {
	counter, err := s.counters.Get(ID)
	if err != nil {
		return errors.WrapNotFound(err, "counter with id %s not found", ID)
	}

	counter.Increment()

	err = s.counters.Save(*counter)
	if err != nil {
		return errors.WrapNotSavable(err, "counter with id %s cannot be updated", ID)
	}

	return nil
}
