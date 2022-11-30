package source_test

import (
	"battleship/pkg/board/source"
	"battleship/pkg/common/test"
	"reflect"
	"testing"
)

func TestBoardFromFile_SetupBoard_reads_from_test_file(t *testing.T) {
	tt := []struct {
		name                string
		filename            string
		expectingError      bool
		expectedCoordinates [][]string
	}{
		{"happy path", "positions_test.txt", false, test.Coordinates},
		{"error while reading", "wrong_file.txt", true, nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := source.New(tc.filename)
			c, err := s.ReadCoordinates()

			if !tc.expectingError {
				if !tc.expectingError && err != nil {
					t.Errorf("unexpected error: %v", err)
				}

				if reflect.DeepEqual(c, tc.expectedCoordinates) {
					t.Error("board have other fields than expected")
				}
			} else {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
			}
		})
	}
}
