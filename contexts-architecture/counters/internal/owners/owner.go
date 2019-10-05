package owners

type Owner struct {
	Email string
}

func New(email string) (*Owner, error) {
	return &Owner{Email: email}, nil
}

type Repository interface {
	GetByEmail(email string) (*Owner, error)
	Save(user Owner) error
}
