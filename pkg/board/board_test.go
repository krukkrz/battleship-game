package board_test

import (
	"battleship/pkg/board"
	"battleship/pkg/common/test"
	"testing"
)

type shot struct {
	coordinates          string
	expectedHit          bool
	expectedAllShipsSunk bool
	currentNumOfShots    int
}

func TestBoard_Shoot(t *testing.T) {
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
			sps := test.BuildShips()
			b := board.New(sps)
			for _, s := range tc.shots {
				isHit, allShipsSunk := b.Shoot(s.coordinates)
				if isHit != s.expectedHit {
					t.Errorf("expected isHit: %v, got: %v for shot: %s", s.expectedHit, isHit, s.coordinates)
				}

				if allShipsSunk != s.expectedAllShipsSunk {
					t.Errorf("expected expectedAllShipsSunk: %v, got: %v for shot: %s", s.expectedAllShipsSunk, allShipsSunk, s.coordinates)
				}

				if b.Shots != s.currentNumOfShots {
					t.Errorf("expected current number of shots: %v, got: %v for shot: %s", s.currentNumOfShots, b.Shots, s.coordinates)
				}
			}
		})
	}
}
