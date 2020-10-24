package inmemory

import (
	"fmt"
	"sync"

	counters "github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal"
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

func (r *usersRepository) Get(ID string) (counters.User, error) {
	user, ok := r.users[ID]
	if !ok {
		return counters.User{}, fmt.Errorf("user id %s: %w", ID, counters.ErrUserNotFound)
	}

	return user, nil
}

func (r *usersRepository) GetByEmail(email string) (counters.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return counters.User{}, fmt.Errorf("user email %s: %w", email, counters.ErrUserNotFound)
}

func (r *usersRepository) Save(user counters.User) error {
	r.users[user.ID] = user
	return nil
}
