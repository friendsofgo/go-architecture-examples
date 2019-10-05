package inmemory

import (
	"sync"

	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/users/internal/counters"

	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/kit/errors"
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

func (r *countersRepository) Get(externalID string) (*counters.Counter, error) {
	counter, ok := r.counters[externalID]
	if !ok {
		return nil, errors.NewNotFound("counter with external id %s not found", externalID)
	}

	return &counter, nil
}

func (r *countersRepository) Save(counter counters.Counter) error {
	r.counters[counter.ExternalID] = counter
	return nil
}
