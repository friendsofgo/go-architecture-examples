package fetching

import (
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/counters/internal/owners"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/kit/errors"
)

type Service interface {
	FetchByEmail(email string) (owners.Owner, error)
}

type service struct {
	owners owners.Repository
}

func NewService(oR owners.Repository) Service {
	return &service{owners: oR}
}

func (s *service) FetchByEmail(email string) (owners.Owner, error) {
	owner, err := s.owners.GetByEmail(email)
	if err != nil {
		return owners.Owner{}, errors.WrapNotFound(err, "owner with email %s not found", email)
	}

	return *owner, nil
}
