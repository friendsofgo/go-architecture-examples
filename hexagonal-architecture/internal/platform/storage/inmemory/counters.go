package inmemory

import (
	"fmt"
	"sync"

	counters "github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal"
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

func (r *countersRepository) Get(ID string) (counters.Counter, error) {
	counter, ok := r.counters[ID]
	if !ok {
		return counters.Counter{}, fmt.Errorf("counter id %s : %w", ID, counters.ErrCounterNotFound)
	}

	return counter, nil
}

func (r *countersRepository) Save(counter counters.Counter) error {
	r.counters[counter.ID] = counter
	return nil
}
