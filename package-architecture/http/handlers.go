package http

import (
	"net/http"

	"github.com/friendsofgo/go-architecture-examples/package-architecture/counters"
	"github.com/friendsofgo/go-architecture-examples/package-architecture/utils"
)

const (
	ok = "ok"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	counterID, err := utils.GetKeyFromHttpRequest(r, "counter")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	counters.CreateCounter(counterID)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ok))
}

func IncrementHandler(w http.ResponseWriter, r *http.Request) {
	counterID, err := utils.GetKeyFromHttpRequest(r, "counter")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	counters.IncrementCounter(counterID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ok))
}
