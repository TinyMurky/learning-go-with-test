package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/build-an-application/repository"
	"example.com/build-an-application/server"
	"example.com/build-an-application/utils"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := repository.NewInMemoryPlayingStore()
	server := server.PlayerServer{
		Store: store,
	}

	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), utils.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), utils.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), utils.NewPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, utils.NewGetScoreRequest(player))

	utils.AssertStatus(t, response.Code, http.StatusOK)
	utils.AssertResponseBody(t, response.Body.String(), "3")

}
