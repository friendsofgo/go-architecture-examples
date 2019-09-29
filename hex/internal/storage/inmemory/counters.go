package inmemory

import (
	"sync"

	counters "github.com/friendsofgo/go-architecture-examples/hex/internal"
	"github.com/friendsofgo/go-architecture-examples/hex/internal/errors"
)

type countersInMemoryRepository struct {
	counters map[string]counters.Counter
}

var (
	countersOnce     sync.Once
	countersInstance *countersInMemoryRepository
)

func NewCountersInMemoryRepository() counters.CounterRepository {
	countersOnce.Do(func() {
		countersInstance = &countersInMemoryRepository{
			counters: make(map[string]counters.Counter),
		}
	})

	return countersInstance
}

func (r *countersInMemoryRepository) Get(ID string) (*counters.Counter, error) {
	counter, ok := r.counters[ID]
	if !ok {
		return nil, errors.NewNotFound("counter with id %s not found", ID)
	}

	return &counter, nil
}

func (r *countersInMemoryRepository) Save(counter counters.Counter) error {
	r.counters[counter.ID] = counter
	return nil
}
