package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/users/internal/users/creating"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/users/internal/users/fetching"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/users/internal/users/storage/inmemory"

	httpserver "github.com/friendsofgo/go-architecture-examples/contexts-architecture/users/cmd/internal/server/http"
)

const (
	ApiHostDefault = "localhost"
	ApiPortDefault = 3000

	ReadTimeoutDefault  = 10 // seconds
	WriteTimeoutDefault = 10 // seconds
)

func main() {

	var (
		usersRepository = inmemory.NewUsersRepository()

		fetchService  = fetching.NewService(usersRepository)
		createService = creating.NewService(usersRepository)

		apiAddress = fmt.Sprintf("%s:%d", ApiHostDefault, ApiPortDefault)
	)

	handler, err := httpserver.MainHandler(fetchService, createService)
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
