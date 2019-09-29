package inmemory

import (
	"sync"

	counters "github.com/friendsofgo/go-architecture-examples/hex/internal"
	"github.com/friendsofgo/go-architecture-examples/hex/internal/errors"
)

type usersRepository struct {
	users map[string]counters.User
}

var (
	usersOnce     sync.Once
	usersInstance *usersRepository
)

func NewUsersRepository() counters.UserRepository {
	usersOnce.Do(func() {
		usersInstance = &usersRepository{
			users: make(map[string]counters.User),
		}
	})

	return usersInstance
}

func (r *usersRepository) Get(ID string) (*counters.User, error) {
	user, ok := r.users[ID]
	if !ok {
		return nil, errors.NewNotFound("user with id %s not found", ID)
	}

	return &user, nil
}

func (r *usersRepository) GetByEmail(email string) (*counters.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, errors.NewNotFound("user with email %s not found", email)
}

func (r *usersRepository) Save(user counters.User) error {
	r.users[user.ID] = user
	return nil
}
