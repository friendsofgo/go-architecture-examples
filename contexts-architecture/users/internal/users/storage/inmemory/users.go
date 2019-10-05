package inmemory

import (
	"sync"

	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/users/internal/users"

	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/kit/errors"
)

type usersRepository struct {
	users map[string]users.User
}

var (
	usersOnce     sync.Once
	usersInstance *usersRepository
)

func NewUsersRepository() users.Repository {
	usersOnce.Do(func() {
		usersInstance = &usersRepository{
			users: make(map[string]users.User),
		}
	})

	return usersInstance
}

func (r *usersRepository) Get(ID string) (*users.User, error) {
	user, ok := r.users[ID]
	if !ok {
		return nil, errors.NewNotFound("user with id %s not found", ID)
	}

	return &user, nil
}

func (r *usersRepository) GetByEmail(email string) (*users.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, errors.NewNotFound("user with email %s not found", email)
}

func (r *usersRepository) Save(user users.User) error {
	r.users[user.ID] = user
	return nil
}
