package creating

import (
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/kit/errors"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/users/internal/users"
)

type Service interface {
	Create(name, email, password string) (users.User, error)
}

type service struct {
	users users.Repository
}

func NewService(uR users.Repository) Service {
	return &service{users: uR}
}

func (s *service) Create(name, email, password string) (users.User, error) {
	newUser, err := users.New(name, email, password)
	if err != nil {
		return users.User{}, err
	}

	err = s.users.Save(*newUser)
	if err != nil {
		return users.User{}, errors.WrapNotSavable(err, "user with email %s cannot be saved", email)
	}

	return *newUser, nil
}
