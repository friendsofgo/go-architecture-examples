package main

var counters = make(map[string]int)

func createCounter(ID string) {
	counters[ID] = 0
}

func incrementCounter(ID string) {
	counters[ID]++
}