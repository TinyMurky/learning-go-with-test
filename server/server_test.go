// Package server run net/http server
package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"example.com/build-an-application/model"
	"example.com/build-an-application/utils"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []model.Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []model.Player {
	return nil
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		winCalls: []string{},
	}

	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := utils.NewGetScoreRequest("Pepper")
		// 用NewRecorder Spy On Response
		response := httptest.NewRecorder()

		// 這是一個ServerHTTP 吃一個ResponseWriter和 *Request
		server.ServeHTTP(response, request)

		// server 回傳int但是還是會變成string
		got := response.Body.String()
		want := "20"

		utils.AssertResponseBody(t, got, want)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := utils.NewGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		utils.AssertResponseBody(t, got, want)
	})

	t.Run("return 404 on missing players", func(t *testing.T) {
		request := utils.NewGetScoreRequest("NotExist")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		utils.AssertStatus(t, got, want)
	})
}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{
		scores:   map[string]int{},
		winCalls: []string{},
	}
	server := NewPlayerServer(store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		playerName := "Pepper"

		request := utils.NewPostWinRequest(playerName)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utils.AssertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != playerName {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], playerName)
		}
	})
}

func TestLeague(t *testing.T) {
	store := &StubPlayerStore{}

	// Implement http handler
	server := NewPlayerServer(store)

	t.Run("it should returns 200 on /league", func(t *testing.T) {
		wantedLeague := []model.Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		assertLeague(t, got, wantedLeague)
		utils.AssertStatus(t, response.Code, http.StatusOK)

		// content type

		if response.Result().Header.Get("content-type") != "application/json" {
			t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
		}
	})
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []model.Player) {
	t.Helper()

	// 單純測試一下可不可以被parse
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return
}

func assertLeague(t testing.TB, got, want []model.Player) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
