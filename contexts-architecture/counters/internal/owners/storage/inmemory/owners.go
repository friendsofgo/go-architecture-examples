package inmemory

import (
	"sync"

	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/counters/internal/owners"

	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/kit/errors"
)

type ownersRepository struct {
	owners map[string]owners.Owner
}

var (
	ownersOnce     sync.Once
	ownersInstance *ownersRepository
)

func NewOwnersRepository() owners.Repository {
	ownersOnce.Do(func() {
		ownersInstance = &ownersRepository{
			owners: make(map[string]owners.Owner),
		}
	})

	return ownersInstance
}

func (r *ownersRepository) GetByEmail(email string) (*owners.Owner, error) {
	for _, owner := range r.owners {
		if owner.Email == email {
			return &owner, nil
		}
	}
	return nil, errors.NewNotFound("owner with email %s not found", email)
}

func (r *ownersRepository) Save(owner owners.Owner) error {
	r.owners[owner.Email] = owner
	return nil
}
