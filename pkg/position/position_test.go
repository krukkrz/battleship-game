package position_test

import (
	"battleship/pkg/position"
	"testing"
)

func TestPosition_Shoot(t *testing.T) {
	tt := []struct {
		name     string
		times    int
		expected bool
	}{
		{"first hit should return false", 1, false},
		{"second hit should return true", 2, true},
		{"third hit should return true", 3, true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := position.New("A1")
			var actual bool
			for i := 0; i < tc.times; i++ {
				actual = p.Shoot()
			}
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
