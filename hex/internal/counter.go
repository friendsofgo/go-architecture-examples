package counters

import (
	"time"

	"github.com/friendsofgo/go-architecture-examples/hex/internal/errors"
	"github.com/friendsofgo/go-architecture-examples/hex/kit/ulid"
)

const (
	minNameLength       = 3
	defaultCounterValue = 0
)

type Counter struct {
	ID         string
	Name       string
	Value      uint
	LastUpdate time.Time
	BelongsTo  string
}

func NewCounter(name, belongsTo string) (*Counter, error) {
	if len(name) < minNameLength {
		return nil, errors.NewWrongInput("counter name %s is too short", name)
	}

	return &Counter{
		ID:         ulid.New(),
		Name:       name,
		Value:      defaultCounterValue,
		LastUpdate: time.Now(),
		BelongsTo:  belongsTo,
	}, nil
}

func (c *Counter) Increment() {
	c.Value++
}

type CounterRepository interface {
	Get(ID string) (*Counter, error)
	Save(counter Counter) error
}
