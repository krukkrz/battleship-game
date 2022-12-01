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

type request struct {
	Player      string `json:"player"`
	Coordinates string `json:"coordinates"`
}

type response struct {
	IsHit        bool `json:"isHit"`
	GameFinished bool `json:"gameFinished"`
}

type round struct {
	req    request
	exp    response
	status int
}

func TestRoute_shoot(t *testing.T) {
	tt := []struct {
		name   string
		rounds []round
	}{
		{"happy path", []round{
			{
				request{"mark", "A1"},
				response{true, false},
				http.StatusOK,
			},
			{
				request{"mark", "A2"},
				response{true, false},
				http.StatusOK,
			},
			{
				request{"mark", "A3"},
				response{true, false},
				http.StatusOK,
			},
			{
				request{"mark", "B1"},
				response{true, false},
				http.StatusOK,
			},
			{
				request{"mark", "B2"},
				response{true, true},
				http.StatusOK,
			},
		}},
		{"missed shot", []round{
			{
				request{"bob", "A1"},
				response{true, false},
				http.StatusOK,
			},
			{
				request{"bob", "A2"},
				response{true, false},
				http.StatusOK,
			},
			{
				request{"bob", "A3"},
				response{true, false},
				http.StatusOK,
			},
			{
				request{"bob", "B1"},
				response{true, false},
				http.StatusOK,
			},
			{
				request{"bob", "C1"},
				response{false, false},
				http.StatusOK,
			},
			{
				request{"bob", "B2"},
				response{true, true},
				http.StatusOK,
			},
		}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			testserver := setupTestServer()
			defer testserver.Close()

			e := httpexpect.New(t, testserver.URL)

			for _, r := range tc.rounds {
				e.POST("/shoot").
					WithJSON(r.req).
					Expect().
					Status(r.status).
					JSON().
					Equal(r.exp)
			}

		})
	}
}

func TestRoute_topTen(t *testing.T) {
	// given
	type topTenResponse struct {
		Name  string `json:"name"`
		Shots int    `json:"shots"`
	}
	rounds := makeWinningRoundsSlice()

	testserver := setupTestServer()
	defer testserver.Close()
	e := httpexpect.New(t, testserver.URL)

	for _, r := range rounds {
		e.POST("/shoot").
			WithJSON(r.req).
			Expect().
			Status(r.status).
			JSON().
			Equal(r.exp)
	}

	// when
	e.GET("/top-ten").
		Expect().
		Status(http.StatusOK).
		JSON().
		Array().
		Contains(topTenResponse{
			Name:  "merry",
			Shots: 5,
		})
}

func setupTestServer() *httptest.Server {
	eng := engine.New(test.Coordinates)
	srv := server.New("", eng)
	return httptest.NewServer(srv.RegisterRoutes())
}

func makeWinningRoundsSlice() []round {
	return []round{
		{
			request{"merry", "A1"},
			response{true, false},
			http.StatusOK,
		},
		{
			request{"merry", "A2"},
			response{true, false},
			http.StatusOK,
		},
		{
			request{"merry", "A3"},
			response{true, false},
			http.StatusOK,
		},
		{
			request{"merry", "B1"},
			response{true, false},
			http.StatusOK,
		},
		{
			request{"merry", "B2"},
			response{true, true},
			http.StatusOK,
		},
	}
}
