package repository

import (
	"fmt"
	"io"

	"example.com/build-an-application/model"
)

// FileSystemPlayerStore a
type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

// GetLeague a
func (f *FileSystemPlayerStore) GetLeague() []model.Player {
	// 讀之前把指標移動回去
	f.database.Seek(0, io.SeekStart) // 回到SeekStart 之後往後讀0 byte
	league, err := model.NewLeague(f.database)
	if err != nil {
		fmt.Println(err)
	}
	return league
}

// GetPlayerScore a
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	league := f.GetLeague()

	for _, player := range league {
		if player.Name == name {
			return player.Wins
		}
	}

	return 0
}
