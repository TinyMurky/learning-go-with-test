package repository

import (
	"io"
	"os"
	"testing"

	"example.com/build-an-application/model"
	"example.com/build-an-application/utils"
)

func TestFileStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		// 這個是用來表示io.Writer
		// string.NewReader 回傳的Reader都可以
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
			]`)

		defer cleanDatabase()

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

		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
			]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		player := "Chris"

		got := store.GetPlayerScore(player)

		want := 33

		utils.AssertDeepEqual(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
			]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		player := "Chris"

		store.RecordWin(player)
		got := store.GetPlayerScore(player)

		want := 34

		utils.AssertDeepEqual(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {

		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
			]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		newPlayer := "Pepper"

		store.RecordWin(newPlayer)

		got := store.GetPlayerScore(newPlayer)
		want := 1
		utils.AssertDeepEqual(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	// 隨機名稱文件，但要自己close + 刪除
	tmpFile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}
