package counters

import (
	"fmt"
	"time"

	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/errors"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/kit/ulid"
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

var (
	ErrCounterNotFound = errors.New("counter not found")
	ErrNameTooShort = errors.New("the receive name is too short")
)

func NewCounter(name, belongsTo string) (*Counter, error) {
	if len(name) < minNameLength {
		return nil, fmt.Errorf("min value is: %d: %w", minNameLength, ErrNameTooShort)
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
