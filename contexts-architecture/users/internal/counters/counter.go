package counters

type Counter struct {
	ExternalID string
	Value      uint
}

func New(externalID string, value uint) *Counter {
	return &Counter{
		ExternalID: externalID,
		Value:      value,
	}
}

func (c *Counter) Update(newValue uint) {
	c.Value = newValue
}

type CounterRepository interface {
	Get(externalID string) (*Counter, error)
	Save(counter Counter) error
}
