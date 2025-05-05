// Package server run net/http server
package server

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore is act like repository
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// PlayerServer Store in stuff
type PlayerServer struct {
	Store PlayerStore
}

// PlayerServer return how many player
func (ps *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ps.processScore(w, r)
	case http.MethodPost:
		ps.processWin(w, r)
	}
}

func (ps *PlayerServer) processScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := ps.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (ps *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	ps.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

// InMemoryPlayerStore implement PlayStore
type InMemoryPlayerStore struct {
	scores map[string]int
}

// GetPlayerScore will return socre by name
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	score, ok := i.scores[name]

	if !ok {
		return 0
	}

	return score
}

// RecordWin will record win
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.scores[name]++
}
