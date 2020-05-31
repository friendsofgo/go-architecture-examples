package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	countershttp "github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/cmd/counters-api/internal/server/http"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/creating"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/fetching"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/incrementing"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/storage/inmemory"
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
		usersRepository    = inmemory.NewUsersRepository()

		fetchService     = fetching.NewFetchService(countersRepository, usersRepository)
		createService    = creating.NewCreateService(countersRepository, usersRepository)
		incrementService = incrementing.NewIncrementService(countersRepository)

		apiAddress = fmt.Sprintf("%s:%d", ApiHostDefault, ApiPortDefault)
	)

	handler, err := countershttp.MainHandler(fetchService, createService, incrementService)
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
