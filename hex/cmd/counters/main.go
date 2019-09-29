package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/friendsofgo/go-architecture-examples/hex/internal/creating"
	"github.com/friendsofgo/go-architecture-examples/hex/internal/fetching"
	"github.com/friendsofgo/go-architecture-examples/hex/internal/incrementing"
	"github.com/friendsofgo/go-architecture-examples/hex/internal/storage/inmemory"

	countershttp "github.com/friendsofgo/go-architecture-examples/hex/internal/server/http"
)

const (
	ApiHostDefault = "localhost"
	ApiPortDefault = 3000

	ReadTimeoutDefault  = 10 // seconds
	WriteTimeoutDefault = 10 // seconds
)

func main() {

	var (
		countersRepository = inmemory.NewCountersInMemoryRepository()
		usersRepository   = inmemory.NewUsersInMemoryRepository()

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

// USERID: 01DNYBKCAFJNTEFRC7S5ZBJX2B
//Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1ycm9ib3R0b0BmcmllbmRzb2Znby5sb2NhbCIsImV4cCI6MTU2OTc1OTIxMSwiaWQiOiIwMUROWUJLQ0FGSk5URUZSQzdTNVpCSlgyQiIsIm5hbWUiOiJNci4gUm9ib3R0byIsIm9yaWdfaWF0IjoxNTY5NzU1NjExfQ.50BqQGBEvfGlPUJ3J0k3gv7t7TVd82EuYac7hicDMu4