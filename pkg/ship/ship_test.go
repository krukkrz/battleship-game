package ship_test

import (
	"battleship/pkg/position"
	"battleship/pkg/ship"
	"testing"
)

type shot struct {
	coordinates   string
	expectedIsHit bool
	expectedSunk  bool
}

func TestShip_Shoot(t *testing.T) {
	tt := []struct {
		name  string
		shots []shot
	}{
		{"all shots are successfull", []shot{
			{"A1", true, false},
			{"A2", true, false},
			{"A3", true, true},
		}},

		{"fourth successfull shot returns same values", []shot{
			{"A1", true, false},
			{"A2", true, false},
			{"A3", true, true},
			{"A3", true, true},
		}},

		{"missed shot returns false in isHit", []shot{
			{"B1", false, false},
		}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ps := buildNewPositions()
			s := ship.New(ps)

			for _, shot := range tc.shots {
				isHit, sunk := s.Shoot(shot.coordinates)
				if isHit != shot.expectedIsHit {
					t.Errorf("expected isHit: %v, got: %v", shot.expectedIsHit, isHit)
				}

				if sunk != shot.expectedSunk {
					t.Errorf("expected sunk: %v, got: %v for shot: %s", shot.expectedSunk, sunk, shot.coordinates)
				}
			}
		})
	}
}

func buildNewPositions() []*position.Position {
	return []*position.Position{
		position.New("A1"),
		position.New("A2"),
		position.New("A3"),
	}
}
