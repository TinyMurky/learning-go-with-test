// Package server run net/http server
package server

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerServer return how many player
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, GetPlayerScore(player))
}

// GetPlayerScore will return socre by name
func GetPlayerScore(name string) string {
	switch name {
	case "Pepper":
		return "20"
	case "Floyd":
		return "10"
	default:
		return ""
	}
}
