package main

import (
	"log"

	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/cmd/http/bootstrap"
)

func main() {
	log.Fatal(bootstrap.Run())
}
