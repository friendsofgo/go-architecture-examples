package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	countershttp "github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/cmd/counters-api/internal/server/http"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/creator"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/fetcher"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/incrementer"
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

		fetcherService     = fetcher.NewService(countersRepository, usersRepository)
		creatorService     = creator.NewService(countersRepository, usersRepository)
		incrementerService = incrementer.NewService(countersRepository)

		apiAddress = fmt.Sprintf("%s:%d", ApiHostDefault, ApiPortDefault)
	)

	handler, err := countershttp.MainHandler(fetcherService, creatorService, incrementerService)
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
