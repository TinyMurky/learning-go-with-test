package utils

import (
	"fmt"
	"net/http"
	"testing"
)

func NewGetScoreRequest(name string) *http.Request {
	url := fmt.Sprintf("/players/%s", name)
	request, _ := http.NewRequest(
		http.MethodGet,
		url,
		nil, // this is body
	)
	return request
}

func NewPostWinRequest(name string) *http.Request {
	url := fmt.Sprintf("/players/%s", name)
	request, _ := http.NewRequest(
		http.MethodPost,
		url,
		nil, // this is body
	)
	return request
}

func AssertResponseBody[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, expect \"%v\", got \"%v\"", want, got)
	}
}

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, expect \"%d\", got \"%d\"", want, got)
	}
}
