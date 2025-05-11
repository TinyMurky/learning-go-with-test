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

func (i *InMemoryPlayerStore) GetLeague() []model.Player {
	return nil
}
