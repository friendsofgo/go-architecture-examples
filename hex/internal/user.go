package counters

type User struct {
	Mail string
	ID string
}

func NewUser(id, mail string) *User {
	return &User{ID: id, Mail: mail}
}

type UserRepository interface {
	Get(ID string) (*User, error)
	Save(user User) error
}
