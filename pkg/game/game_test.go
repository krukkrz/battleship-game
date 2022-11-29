package game_test

import (
	"battleship/pkg/common/test"
	"battleship/pkg/game"
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
			b := test.BuildBoard()
			g := game.New("mark", &b)

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
