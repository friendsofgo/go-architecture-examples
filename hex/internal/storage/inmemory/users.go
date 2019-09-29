package inmemory

import (
	"sync"

	counters "github.com/friendsofgo/go-architecture-examples/hex/internal"
	"github.com/friendsofgo/go-architecture-examples/hex/internal/errors"
)

type usersInMemoryRepository struct {
	users map[string]counters.User
}

var (
	usersOnce     sync.Once
	usersInstance *usersInMemoryRepository
)

func NewUsersInMemoryRepository() counters.UserRepository {
	usersOnce.Do(func() {
		usersInstance = &usersInMemoryRepository{
			users: make(map[string]counters.User),
		}
	})

	return usersInstance
}

func (r *usersInMemoryRepository) Get(ID string) (*counters.User, error) {
	user, ok := r.users[ID]
	if !ok {
		return nil, errors.NewNotFound("user with id %s not found", ID)
	}

	return &user, nil
}

func (r *usersInMemoryRepository) Save(user counters.User) error {
	r.users[user.ID] = user
	return nil
}
