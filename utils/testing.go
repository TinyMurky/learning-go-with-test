package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"example.com/build-an-application/model"
)

// NewGetScoreRequest a
func NewGetScoreRequest(name string) *http.Request {
	url := fmt.Sprintf("/players/%s", name)
	request, _ := http.NewRequest(
		http.MethodGet,
		url,
		nil, // this is body
	)
	return request
}

// NewPostWinRequest a
func NewPostWinRequest(name string) *http.Request {
	url := fmt.Sprintf("/players/%s", name)
	request, _ := http.NewRequest(
		http.MethodPost,
		url,
		nil, // this is body
	)
	return request
}

// AssertResponseBody a
func AssertResponseBody[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, expect \"%v\", got \"%v\"", want, got)
	}
}

// AssertStatus a
func AssertStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, expect \"%d\", got \"%d\"", want, got)
	}
}

// NewLeagueRequest a
func NewLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

// GetLeagueFromResponse a
func GetLeagueFromResponse(t testing.TB, body io.Reader) (league []model.Player) {
	t.Helper()

	// 單純測試一下可不可以被parse
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return
}

// AssertLeague a
func AssertLeague(t testing.TB, got, want []model.Player) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// AssertContentType a
func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()

	if response.Header().Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Header())
	}
}

// AssertDeepEqual a
func AssertDeepEqual[T any](t testing.TB, got, want T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expect %+v, got %+v", want, got)

	}
}
