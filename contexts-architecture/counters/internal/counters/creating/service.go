package creating

import (
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/counters/internal/counters"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/kit/errors"
)

type Service interface {
	Create(name, belongsTo string) (counters.Counter, error)
}

type service struct {
	counters counters.Repository
}

func NewService(cR counters.Repository) Service {
	return &service{counters: cR}
}

func (s *service) Create(name, belongsTo string) (counters.Counter, error) {
	newCounter, err := counters.New(name, belongsTo)
	if err != nil {
		return counters.Counter{}, err
	}

	err = s.counters.Save(*newCounter)
	if err != nil {
		return counters.Counter{}, errors.WrapNotSavable(err, "counter with name %s cannot be saved", name)
	}

	return *newCounter, nil
}
