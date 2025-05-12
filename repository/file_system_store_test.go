package repository

import (
	"strings"
	"testing"

	"example.com/build-an-application/model"
	"example.com/build-an-application/utils"
)

func TestFileStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		// 這個是用來表示io.Writer

		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
			]`)

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []model.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}

		utils.AssertLeague(t, got, want)
	})
}
