package game_test

import (
	"battleship/pkg/common/test"
	"battleship/pkg/game"
	"testing"
)

type shot struct {
	coordinates        string
	expectedHit        bool
	expectedFinished   bool
	expectedNumOfShots int
}

func TestGame_Shoot(t *testing.T) {
	tt := []struct {
		name  string
		shots []shot
	}{
		{"all shots successfull", []shot{
			{"A1", true, false, 1},
			{"A2", true, false, 2},
			{"A3", true, false, 3},
			{"B1", true, false, 4},
			{"B2", true, true, 5},
		}},
		{"successfull game with missed shot", []shot{
			{"A1", true, false, 1},
			{"A2", true, false, 2},
			{"C2", false, false, 3},
			{"A3", true, false, 4},
			{"B1", true, false, 5},
			{"B2", true, true, 6},
		}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			b := test.BuildBoard()
			g := game.New(b)

			for _, s := range tc.shots {
				isHit, finished := g.Shoot(s.coordinates)
				if isHit != s.expectedHit {
					t.Errorf("expected isHit: %v, got: %v for shot: %s", s.expectedHit, isHit, s.coordinates)
				}

				if finished != s.expectedFinished {
					t.Errorf("expected expectedAllShipsSunk: %v, got: %v for shot: %s", s.expectedFinished, finished, s.coordinates)
				}

				if g.Shots != s.expectedNumOfShots {
					t.Errorf("expected current number of shots: %v, got: %v for shot: %s", s.expectedNumOfShots, g.Shots, s.coordinates)
				}
			}
		})
	}
}
