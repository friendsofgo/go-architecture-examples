package inmemory

import (
	"sync"

	counters "github.com/friendsofgo/go-architecture-examples/hex/internal"
	"github.com/friendsofgo/go-architecture-examples/hex/internal/errors"
)

type countersRepository struct {
	counters map[string]counters.Counter
}

var (
	countersOnce     sync.Once
	countersInstance *countersRepository
)

func NewCountersRepository() counters.CounterRepository {
	countersOnce.Do(func() {
		countersInstance = &countersRepository{
			counters: make(map[string]counters.Counter),
		}
	})

	return countersInstance
}

func (r *countersRepository) Get(ID string) (*counters.Counter, error) {
	counter, ok := r.counters[ID]
	if !ok {
		return nil, errors.NewNotFound("counter with id %s not found", ID)
	}

	return &counter, nil
}

func (r *countersRepository) Save(counter counters.Counter) error {
	r.counters[counter.ID] = counter
	return nil
}
