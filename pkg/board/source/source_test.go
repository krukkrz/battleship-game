package source_test

import (
	"battleship/pkg/board/source"
	"testing"
)

type shot struct {
	coordinates   string
	expectedIsHit bool
	expectedSunk  bool
}

func TestBoardFromFile_SetupBoard_reads_from_test_file(t *testing.T) {
	tt := []struct {
		name          string
		filename      string
		expectedError error
		shots         []shot
	}{
		{"happy path", "positions_test.txt", nil, []shot{
			{"A1", true, false},
			{"A2", true, false},
			{"A3", true, false},
			{"B1", true, false},
			{"B2", true, true},
		}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := source.New(tc.filename)
			b, err := s.SetupBoard()

			if err != tc.expectedError {
				t.Errorf("expected error: %v, got: %v", tc.expectedError, err)
			}

			for _, s := range tc.shots {
				isHit, sunk := b.Shoot(s.coordinates)
				if isHit != s.expectedIsHit || sunk != s.expectedSunk {
					t.Error("board have other fields than expected")
				}
			}
		})
	}
}
