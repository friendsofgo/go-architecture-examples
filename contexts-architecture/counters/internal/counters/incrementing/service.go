package incrementing

import (
	counters "github.com/friendsofgo/go-architecture-examples/contexts-architecture/counters/internal/counters"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/kit/errors"
)

type Service interface {
	Increment(ID string) error
}

type service struct {
	counters counters.Repository
}

func NewService(cR counters.Repository) Service {
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
