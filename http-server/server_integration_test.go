package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordWinsAndRetrieveThem(t *testing.T) {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	player := "Pepper"

	for range 3 {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertResponseHeader(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")

}
