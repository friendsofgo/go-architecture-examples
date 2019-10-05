package creating

import (
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/counters/internal/owners"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/kit/errors"
)

type Service interface {
	Create(email string) (owners.Owner, error)
}

type service struct {
	owners owners.Repository
}

func NewService(oR owners.Repository) Service {
	return &service{owners: oR}
}

func (s *service) Create(email string) (owners.Owner, error) {
	newOwner, err := owners.New(email)
	if err != nil {
		return owners.Owner{}, err
	}

	err = s.owners.Save(*newOwner)
	if err != nil {
		return owners.Owner{}, errors.WrapNotSavable(err, "owner with email %s cannot be saved", email)
	}

	return *newOwner, nil
}
