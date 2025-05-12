// Package server run net/http server
package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"example.com/build-an-application/model"
)

const jsonContentType = "application/json"

// PlayerStore is act like repository
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []model.Player
}

// PlayerServer Store in stuff
type PlayerServer struct {
	Store PlayerStore
	http.Handler
}

// NewPlayerServer create new server
func NewPlayerServer(store PlayerStore) *PlayerServer {

	server := &PlayerServer{
		Store: store,
	}
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(server.leagueHandler))

	// 注意這裡是 /players/
	router.Handle("/players/", http.HandlerFunc(server.playerHandler))

	server.Handler = router
	return server
}

func (ps *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", jsonContentType) // 先set header
	w.WriteHeader(http.StatusOK)                    // 再set status
	json.NewEncoder(w).Encode(ps.Store.GetLeague()) // 最後寫body
}

func (ps *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodGet:
		ps.showScore(w, player)
	case http.MethodPost:
		ps.processWin(w, player)
	}
}
func (ps *PlayerServer) showScore(w http.ResponseWriter, player string) {

	score := ps.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (ps *PlayerServer) processWin(w http.ResponseWriter, player string) {
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
