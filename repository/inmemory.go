package repository

import (
	"sync"

	"example.com/build-an-application/model"
)

// InMemoryPlayerStore a
type InMemoryPlayerStore struct {
	store map[string]int
	mux   sync.Mutex
}

// NewInMemoryPlayingStore a
func NewInMemoryPlayingStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: map[string]int{},
	}
}

// GetPlayerScore aaa
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mux.Lock()
	defer i.mux.Unlock()

	return i.store[name]
}

// RecordWin aa
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mux.Lock()
	defer i.mux.Unlock()

	i.store[name]++
}

// GetLeague a
func (i *InMemoryPlayerStore) GetLeague() []model.Player {
	league := make([]model.Player, 0, len(i.store))

	for name, wins := range i.store {
		league = append(league, model.Player{
			Name: name,
			Wins: wins,
		})
	}

	return league
}
