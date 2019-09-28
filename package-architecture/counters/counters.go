package counters

var counters = make(map[string]int)

func CreateCounter(ID string) {
	counters[ID] = 0
}

func IncrementCounter(ID string) {
	counters[ID]++
}