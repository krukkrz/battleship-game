package game_test

import (
	"battleship/pkg/board"
	"battleship/pkg/game"
	"battleship/pkg/position"
	"battleship/pkg/ship"
	"testing"
)

type shot struct {
	coordinates      string
	expectedHit      bool
	expectedFinished bool
}

func TestGame_Shoot(t *testing.T) {
	tt := []struct {
		name  string
		shots []shot
	}{
		{"all shots successfull", []shot{
			{"A1", true, false},
			{"A2", true, false},
			{"A3", true, false},
			{"B1", true, false},
			{"B2", true, true},
		}},
		{"successfull game with missed shot", []shot{
			{"A1", true, false},
			{"A2", true, false},
			{"C2", false, false},
			{"A3", true, false},
			{"B1", true, false},
			{"B2", true, true},
		}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			b := buildBoard()
			g := game.New("mark", b)

			for _, s := range tc.shots {
				isHit, finished := g.Shoot(s.coordinates)
				if isHit != s.expectedHit {
					t.Errorf("expected isHit: %v, got: %v for shot: %s", s.expectedHit, isHit, s.coordinates)
				}

				if finished != s.expectedFinished {
					t.Errorf("expected expectedAllShipsSunk: %v, got: %v for shot: %s", s.expectedFinished, finished, s.coordinates)
				}
			}
		})
	}
}

func buildBoard() *board.Board {
	sps := buildShips()
	return board.New(sps)
}

func buildShips() []*ship.Ship {
	return []*ship.Ship{
		buildShip("A1", "A2", "A3"),
		buildShip("B1", "B2"),
	}
}

func buildShip(coordinates ...string) *ship.Ship {
	return ship.New(buildPositions(coordinates...))
}

func buildPositions(coordinates ...string) []*position.Position {
	var p []*position.Position
	for _, c := range coordinates {
		p = append(p, position.New(c))
	}
	return p
}
