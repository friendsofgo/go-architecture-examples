package counters

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/kit/ulid"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrCreatingUser = errors.New("err creating user")
)

type User struct {
	ID             string
	Name           string
	Email          string
	HashedPassword string
}

func NewUser(name, email, password string) (User, error) {
	u := User{
		ID:    ulid.New(),
		Name:  name,
		Email: email,
	}

	err := u.HashPassword(password)
	if err != nil {
		return User{}, fmt.Errorf("password cannot be hashed correctly: %w", ErrCreatingUser)
	}

	return u, nil
}

func (u *User) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.HashedPassword = string(hash)
	return nil
}

type UserRepository interface {
	Get(ID string) (User, error)
	GetByEmail(email string) (User, error)
	Save(user User) error
}
