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
		// string.NewReader 回傳的Reader都可以
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

		// 確保可以重複讀
		got2 := store.GetLeague()
		utils.AssertLeague(t, got2, want)
	})

	t.Run("get player score", func(t *testing.T) {

		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
			]`)

		store := FileSystemPlayerStore{database}

		player := "Chris"

		got := store.GetPlayerScore(player)

		want := 33

		utils.AssertDeepEqual(t, got, want)
	})
}
