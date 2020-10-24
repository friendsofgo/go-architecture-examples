package counters

import (
	"errors"
	"fmt"
	"time"

	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/kit/ulid"
)

const (
	minNameLength       = 3
	defaultCounterValue = 0
)

var (
	ErrCounterNotFound = errors.New("counter not found")
	ErrNameTooShort    = errors.New("counter name is too short")
)

type Counter struct {
	ID         string
	Name       string
	Value      uint
	LastUpdate time.Time
	BelongsTo  string
}

func NewCounter(name, belongsTo string) (Counter, error) {
	if len(name) < minNameLength {
		return Counter{}, fmt.Errorf("min value is: %d: %w", minNameLength, ErrNameTooShort)
	}

	return Counter{
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
	Get(ID string) (Counter, error)
	Save(counter Counter) error
}
