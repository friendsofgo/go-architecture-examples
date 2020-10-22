package main

import (
	"log"

	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/platform/bootstrap"
)

func main() {
	log.Fatal(bootstrap.Run())
}
