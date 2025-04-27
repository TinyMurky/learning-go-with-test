// Package main is where server run
package main

import (
	"log"
	"net/http"

	"example.com/build-an-application/server"
)

func main() {
	// 直接把function包成interface
	handler := http.HandlerFunc(server.PlayerServer)

	err := http.ListenAndServe(":5000", handler)

	if err != nil {
		log.Fatal(err)
	}
}
