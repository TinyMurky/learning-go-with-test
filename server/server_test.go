// Package server run net/http server
package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		// 用NewRecorder Spy On Response
		response := httptest.NewRecorder()

		// 這是一個ServerHTTP 吃一個ResponseWriter和 *Request
		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		assertResponseBody(t, got, want)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "10"

		assertResponseBody(t, got, want)
	})
}

func newGetScoreRequest(name string) *http.Request {
	url := fmt.Sprintf("/players/%s", name)
	request, _ := http.NewRequest(
		http.MethodGet,
		url,
		nil, // this is body
	)
	return request
}

func assertResponseBody[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, expect \"%v\", got \"%v\"", want, got)
	}
}
