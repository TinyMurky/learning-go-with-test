package repository

import (
	"encoding/json"
	"fmt"
	"io"

	"example.com/build-an-application/model"
)

// FileSystemPlayerStore a
type FileSystemPlayerStore struct {
	// Read、Write 和 Seek 操作共用同一個「指標」位置。
	database io.ReadWriteSeeker
}

// GetLeague a
func (f *FileSystemPlayerStore) GetLeague() model.League {
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

// RecordWin a
func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()

	// player 是pointer
	player := league.Find(name)
	player.Wins++

	// 移到開頭
	f.database.Seek(0, io.SeekStart) // Read 過也要重製才能seek
	json.NewEncoder(f.database).Encode(league)
}
