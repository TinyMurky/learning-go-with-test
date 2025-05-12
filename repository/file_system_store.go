package repository

import (
	"fmt"
	"io"

	"example.com/build-an-application/model"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func (f *FileSystemPlayerStore) GetLeague() []model.Player {
	league, err := model.NewLeague(f.database)
	if err != nil {
		fmt.Println(err)
	}
	return league
}
