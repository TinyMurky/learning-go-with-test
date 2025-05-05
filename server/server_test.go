// Package server run net/http server
package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/build-an-application/utils"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		winCalls: []string{},
	}

	server := &PlayerServer{
		Store: &store,
	}

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

	server := &PlayerServer{
		Store: store,
	}

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
