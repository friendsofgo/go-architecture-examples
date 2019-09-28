package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("initialize counters")
	counters = make(map[string]int)

	http.HandleFunc("/counters/increment", incrementHandler)
	http.HandleFunc("/counters/create", createHandler)

	fmt.Println("Server tap on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}