package fetcher

import (
	counters "github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/errors"
)

type Service interface {
	FetchCounterByID(id string) (counters.Counter, error)
	FetchUserByID(id string) (counters.User, error)
	FetchUserByEmail(email string) (counters.User, error)
}

type service struct {
	counters counters.CounterRepository
	users    counters.UserRepository
}

func NewService(cR counters.CounterRepository, uR counters.UserRepository) Service {
	return &service{counters: cR, users: uR}
}

func (s *service) FetchCounterByID(id string) (counters.Counter, error) {
	counter, err := s.counters.Get(id)
	if err != nil {
		return counters.Counter{}, errors.WrapNotFound(err, "counter with id %s not found", id)
	}

	return *counter, nil
}

func (s *service) FetchUserByEmail(email string) (counters.User, error) {
	user, err := s.users.GetByEmail(email)
	if err != nil {
		return counters.User{}, errors.WrapNotFound(err, "user with id %s not found", email)
	}

	return *user, nil
}

func (s *service) FetchUserByID(id string) (counters.User, error) {
	user, err := s.users.Get(id)
	if err != nil {
		return counters.User{}, errors.WrapNotFound(err, "user with id %s not found", id)
	}

	return *user, nil

}
