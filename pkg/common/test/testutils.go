package test

import (
	"battleship/pkg/board"
	"battleship/pkg/ship"
)

var Coordinates = [][]string{
	[]string{"A1", "A2", "A3"},
	[]string{"B1", "B2"},
}

func BuildBoard() *board.Board {
	sps := BuildShips()
	return board.New(sps)
}

func BuildShips() []*ship.Ship {
	return []*ship.Ship{
		ship.New("A1", "A2", "A3"),
		ship.New("B1", "B2"),
	}
}
