package creating

import (
	"fmt"

	counters "github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal"
)

type Service interface {
	CreateCounter(name, belongsTo string) (counters.Counter, error)
	CreateUser(name, email, password string) (counters.User, error)
}

type service struct {
	counters counters.CounterRepository
	users    counters.UserRepository
}

func NewService(cR counters.CounterRepository, uR counters.UserRepository) Service {
	return &service{counters: cR, users: uR}
}

func (s *service) CreateCounter(name, belongsTo string) (counters.Counter, error) {
	newCounter, err := counters.NewCounter(name, belongsTo)
	if err != nil {
		return counters.Counter{}, err
	}

	err = s.counters.Save(*newCounter)
	if err != nil {
		return counters.Counter{}, fmt.Errorf( "counter with name %s cannot be saved", name)
	}

	return *newCounter, nil
}

func (s *service) CreateUser(name, email, password string) (counters.User, error) {
	newUser, err := counters.NewUser(name, email, password)
	if err != nil {
		return counters.User{}, err
	}

	err = s.users.Save(*newUser)
	if err != nil {
		return counters.User{}, fmt.Errorf("user with email %s cannot be saved", email)
	}

	return *newUser, nil
}
