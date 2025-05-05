// Package main is where server run
package main

import (
	"log"
	"net/http"

	"example.com/build-an-application/repository"
	"example.com/build-an-application/server"
)

func main() {

	// 直接把function包成interface
	server := &server.PlayerServer{
		Store: repository.NewInMemoryPlayingStore(),
	}

	err := http.ListenAndServe(":5000", server)

	if err != nil {
		log.Fatal(err)
	}
}
