package model

import (
	"encoding/json"
	"fmt"
	"io"
)

// League a
type League []Player

// NewLeague a
func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		err = fmt.Errorf("problem parsing league: %w", err)
	}

	return league, err
}

func (l League) Find(name string) *Player {
	// range 會複製值
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}

	return nil
}
