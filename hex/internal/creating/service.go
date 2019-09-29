package creating

import (
	counters "github.com/friendsofgo/go-architecture-examples/hex/internal"
	"github.com/friendsofgo/go-architecture-examples/hex/internal/errors"
)

type Service interface {
	CreateCounter(name, belongsTo string) (counters.Counter, error)
	CreateUser(id, mail string) (counters.User, error)
}

type service struct {
	counters  counters.CounterRepository
	users     counters.UserRepository
}

func NewCreateService(cR counters.CounterRepository, uR counters.UserRepository) Service {
	return &service{counters: cR, users: uR}
}

func (s *service) CreateCounter(name, belongsTo string) (counters.Counter, error) {
	newCounter, err := counters.NewCounter(name, belongsTo)
	if err != nil {
		return counters.Counter{}, err
	}

	err = s.counters.Save(*newCounter)
	if err != nil {
		return counters.Counter{}, errors.WrapNotSavable(err, "counter with name %s cannot be saved", name)
	}

	return *newCounter, nil
}

func (s *service) CreateUser(id, mail string) (counters.User, error) {
	newUser := counters.NewUser(id, mail)
	err := s.users.Save(*newUser)
	if err != nil {
		return counters.User{}, errors.WrapNotSavable(err, "user with id %s cannot be saved", id)
	}
	return *newUser, nil
}
