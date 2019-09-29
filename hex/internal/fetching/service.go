package fetching

import (
	counters "github.com/friendsofgo/go-architecture-examples/hex/internal"
	"github.com/friendsofgo/go-architecture-examples/hex/internal/errors"
)

type Service interface {
	Fetch(ID string) (counters.Counter, error)
}

type service struct {
	counters counters.CounterRepository
}

func NewFetchService(cR counters.CounterRepository) Service {
	return &service{counters:cR}
}

func (s *service) Fetch(ID string) (counters.Counter, error) {
	counter, err := s.counters.Get(ID)
	if err != nil {
		return counters.Counter{}, errors.WrapNotFound(err, "counter with id %s not found", ID)
	}

	return *counter, nil
}
