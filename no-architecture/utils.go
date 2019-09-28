package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func getKeyFromHttpRequest(r *http.Request, key string) (string, error) {
	keys, ok := r.URL.Query()[key]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return "", errors.New(fmt.Sprintf("%s param is missing", key))
	}

	return keys[0], nil
}
