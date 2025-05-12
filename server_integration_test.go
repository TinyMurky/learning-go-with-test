package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/build-an-application/model"
	"example.com/build-an-application/repository"
	"example.com/build-an-application/server"
	"example.com/build-an-application/utils"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := repository.NewInMemoryPlayingStore()
	server := server.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), utils.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), utils.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), utils.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, utils.NewGetScoreRequest(player))

		utils.AssertStatus(t, response.Code, http.StatusOK)
		utils.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {

		response := httptest.NewRecorder()
		server.ServeHTTP(response, utils.NewLeagueRequest())

		got := utils.GetLeagueFromResponse(t, response.Body)

		expect := []model.Player{
			{Name: "Pepper", Wins: 3},
		}

		utils.AssertLeague(t, got, expect)
	})

}
