package incrementing

import (
	"fmt"

	counters "github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal"
)

type Service interface {
	Increment(ID string) error
}

type service struct {
	counters counters.CounterRepository
}

func NewService(cR counters.CounterRepository) Service {
	return &service{counters: cR}
}

func (s *service) Increment(ID string) error {
	counter, err := s.counters.Get(ID)
	if err != nil {
		return err
	}

	counter.Increment()

	err = s.counters.Save(*counter)
	if err != nil {
		return fmt.Errorf("counter with id %s cannot be updated", ID)
	}

	return nil
}
