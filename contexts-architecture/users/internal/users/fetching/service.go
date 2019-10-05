package fetching

import (
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/kit/errors"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/users/internal/users"
)

type Service interface {
	FetchByID(id string) (users.User, error)
	FetchByEmail(email string) (users.User, error)
}

type service struct {
	users users.Repository
}

func NewService(uR users.Repository) Service {
	return &service{users: uR}
}

func (s *service) FetchByEmail(email string) (users.User, error) {
	user, err := s.users.GetByEmail(email)
	if err != nil {
		return users.User{}, errors.WrapNotFound(err, "user with email %s not found", email)
	}

	return *user, nil
}

func (s *service) FetchByID(id string) (users.User, error) {
	user, err := s.users.Get(id)
	if err != nil {
		return users.User{}, errors.WrapNotFound(err, "user with id %s not found", id)
	}

	return *user, nil

}
