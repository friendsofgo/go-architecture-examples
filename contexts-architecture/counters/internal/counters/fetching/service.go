package fetching

import (
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/counters/internal/counters"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/kit/errors"
)

type Service interface {
	FetchByID(id string) (counters.Counter, error)
}

type service struct {
	counters counters.Repository
}

func NewService(cR counters.Repository) Service {
	return &service{counters: cR}
}

func (s *service) FetchByID(id string) (counters.Counter, error) {
	counter, err := s.counters.Get(id)
	if err != nil {
		return counters.Counter{}, errors.WrapNotFound(err, "counter with id %s not found", id)
	}

	return *counter, nil
}
