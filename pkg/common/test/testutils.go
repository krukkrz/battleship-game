package test

import (
	"battleship/pkg/board"
	"battleship/pkg/position"
	"battleship/pkg/ship"
)

func BuildBoard() board.Board {
	sps := BuildShips()
	return *board.New(sps)
}

func BuildShips() []*ship.Ship {
	return []*ship.Ship{
		buildShip("A1", "A2", "A3"),
		buildShip("B1", "B2"),
	}
}

func buildShip(coordinates ...string) *ship.Ship {
	return ship.New(BuildPositions(coordinates...))
}

func BuildPositions(coordinates ...string) []*position.Position {
	var p []*position.Position
	for _, c := range coordinates {
		p = append(p, position.New(c))
	}
	return p
}
