package creating

import (
	counters "github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal"
)

type Service interface {
	CreateCounter(name, belongsTo string) (counters.Counter, error)
	CreateUser(name, email, password string) (counters.User, error)
}

type DefaultService struct {
	counters counters.CounterRepository
	users    counters.UserRepository
}

func NewService(cR counters.CounterRepository, uR counters.UserRepository) DefaultService {
	return DefaultService{counters: cR, users: uR}
}

func (s DefaultService) CreateCounter(name, belongsTo string) (counters.Counter, error) {
	newCounter, err := counters.NewCounter(name, belongsTo)
	if err != nil {
		return counters.Counter{}, err
	}

	err = s.counters.Save(newCounter)
	if err != nil {
		return counters.Counter{}, err
	}

	return newCounter, nil
}

func (s DefaultService) CreateUser(name, email, password string) (counters.User, error) {
	newUser, err := counters.NewUser(name, email, password)
	if err != nil {
		return counters.User{}, err
	}

	err = s.users.Save(newUser)
	if err != nil {
		return counters.User{}, err
	}

	return newUser, nil
}
