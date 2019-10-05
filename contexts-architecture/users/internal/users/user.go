package users

import (
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/kit/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/kit/ulid"
)

type User struct {
	ID             string
	Name           string
	Email          string
	HashedPassword string
}

func New(name, email, password string) (*User, error) {
	// validations about your user creation here...

	u := User{
		ID:    ulid.New(),
		Name:  name,
		Email: email,
	}

	err := u.HashPassword(password)
	if err != nil {
		return nil, errors.WrapWrongInput(err, "user password %s cannot be hashed", password)
	}

	return &u, nil
}

func (u *User) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.HashedPassword = string(hash)
	return nil
}

type Repository interface {
	Get(ID string) (*User, error)
	GetByEmail(email string) (*User, error)
	Save(user User) error
}
