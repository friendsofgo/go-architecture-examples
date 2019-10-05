package creating

import (
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/kit/errors"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/users/internal/counters"
)

type Service interface {
	Create(externalID string, value uint) (counters.Counter, error)
}

type service struct {
	counters counters.CounterRepository
}

func NewService(cR counters.CounterRepository) Service {
	return &service{counters: cR}
}

func (s *service) Create(externalID string, value uint) (counters.Counter, error) {
	newCounter := counters.New(externalID, value)

	err := s.counters.Save(*newCounter)
	if err != nil {
		return counters.Counter{}, errors.WrapNotSavable(err, "counter with external id %s cannot be saved", externalID)
	}

	return *newCounter, nil
}
