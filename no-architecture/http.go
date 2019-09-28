package main

import (
	"net/http"
)

const (
	ok = "ok"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	counterID, err := getKeyFromHttpRequest(r, "counter")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	createCounter(counterID)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ok))
}

func incrementHandler(w http.ResponseWriter, r *http.Request) {
	counterID, err := getKeyFromHttpRequest(r, "counter")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	incrementCounter(counterID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ok))
}
