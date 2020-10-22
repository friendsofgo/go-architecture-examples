package fetching

import (
	counters "github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal"
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
	return service{counters: cR, users: uR}
}

func (s *service) FetchCounterByID(id string) (counters.Counter, error) {
	counter, err := s.counters.Get(id)
	if err != nil {
		return counters.Counter{}, err
	}

	return *counter, nil
}

func (s *service) FetchUserByEmail(email string) (counters.User, error) {
	user, err := s.users.GetByEmail(email)
	if err != nil {
		return counters.User{}, err
	}

	return *user, nil
}

func (s *service) FetchUserByID(id string) (counters.User, error) {
	user, err := s.users.Get(id)
	if err != nil {
		return counters.User{}, err
	}

	return *user, nil

}
