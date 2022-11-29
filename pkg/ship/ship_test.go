package ship_test

import (
	"battleship/pkg/ship"
	"testing"
)

type shot struct {
	coordinates  string
	expectedSunk bool
}

func TestShip_Shoot(t *testing.T) {
	tt := []struct {
		name  string
		shots []shot
	}{
		{"all shots are successfull", []shot{
			{"A1", false},
			{"A2", false},
			{"A3", true},
		}},

		{"fourth successfull shot returns same values", []shot{
			{"A1", false},
			{"A2", false},
			{"A3", true},
			{"A3", true},
		}},

		{"missed shot returns false in isHit", []shot{
			{"B1", false},
		}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := ship.New("A1", "A2", "A3")

			for _, shot := range tc.shots {
				sunk := s.MarkHit(shot.coordinates)

				if sunk != shot.expectedSunk {
					t.Errorf("expected sunk: %v, got: %v for shot: %s", shot.expectedSunk, sunk, shot.coordinates)
				}
			}
		})
	}
}
