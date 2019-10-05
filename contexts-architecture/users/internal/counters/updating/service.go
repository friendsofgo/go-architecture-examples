package updating

import (
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/kit/errors"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/users/internal/counters"
)

type Service interface {
	Update(externalID string, value uint) error
}

type service struct {
	counters counters.CounterRepository
}

func NewService(cR counters.CounterRepository) Service {
	return &service{counters: cR}
}

func (s *service) Update(externalID string, value uint) error {
	counter, err := s.counters.Get(externalID)
	if err != nil {
		return errors.WrapNotFound(err, "counter with external id %s not found", externalID)
	}

	counter.Update(value)

	err = s.counters.Save(*counter)
	if err != nil {
		return errors.WrapNotSavable(err, "counter with id %s cannot be updated", externalID)
	}

	return nil
}
