package server_test

import (
	"battleship/pkg/common/test"
	"battleship/pkg/engine"
	"battleship/pkg/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestRoute_shoot(t *testing.T) {
	// TODO make it table driven test and go through main scenarios
	req := struct {
		Player      string `json:"player"`
		Coordinates string `json:"coordinates"`
	}{"mark", "A1"}

	expectedResponse := struct {
		IsHit        bool `json:"isHit"`
		GameFinished bool `json:"gameFinished"`
	}{true, false}

	testserver := setupTestServer()
	defer testserver.Close()

	e := httpexpect.New(t, testserver.URL)
	e.POST("/shoot").
		WithJSON(req).
		Expect().
		Status(http.StatusOK).
		JSON().
		Equal(expectedResponse)
}

// TODO make a table driven test expecting bad scenarios

func setupTestServer() *httptest.Server {
	eng := engine.New(test.Coordinates)
	srv := server.New("", eng)
	return httptest.NewServer(srv.RegisterRoutes())
}
