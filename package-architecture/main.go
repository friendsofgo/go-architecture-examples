package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/friendsofgo/go-architecture-examples/package-architecture/http"
)

func main() {
	http.HandleFunc("/counters/increment", handler.IncrementHandler)
	http.HandleFunc("/counters/create", handler.CreateHandler)

	fmt.Println("Server tap on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
