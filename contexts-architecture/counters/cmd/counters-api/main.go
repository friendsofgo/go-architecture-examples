package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/counters/internal/counters/creating"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/counters/internal/counters/fetching"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/counters/internal/counters/incrementing"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/counters/internal/counters/storage/inmemory"

	httpserver "github.com/friendsofgo/go-architecture-examples/contexts-architecture/counters/cmd/counters-api/internal/server/http"
)

const (
	ApiHostDefault = "localhost"
	ApiPortDefault = 3000

	ReadTimeoutDefault  = 10 // seconds
	WriteTimeoutDefault = 10 // seconds
)

func main() {

	var (
		countersRepository = inmemory.NewCountersRepository()

		fetchService     = fetching.NewService(countersRepository)
		createService    = creating.NewService(countersRepository)
		incrementService = incrementing.NewService(countersRepository)

		apiAddress = fmt.Sprintf("%s:%d", ApiHostDefault, ApiPortDefault)
	)

	handler, err := httpserver.MainHandler(fetchService, createService, incrementService)
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:         apiAddress,
		Handler:      handler,
		ReadTimeout:  ReadTimeoutDefault * time.Second,
		WriteTimeout: WriteTimeoutDefault * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
