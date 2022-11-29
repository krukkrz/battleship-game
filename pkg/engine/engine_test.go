package engine_test

import (
	"battleship/pkg/board"
	"battleship/pkg/engine"
	"battleship/pkg/position"
	"battleship/pkg/ship"
	"testing"
)

func TestBattleShipGameEngine_New(t *testing.T) {
	// should not create second instnce if there is one already
	b := buildBoard()
	e1 := engine.New(b)
	e2 := engine.New(b)
	if e1 != e2 {
		t.Error("BattleShipGameEngine should be singleton but is not")
	}
}

func TestBattleShipGameEngine_Shoot(t *testing.T) {
}

func TestBattleShipGameEngine_TopTen(t *testing.T) {
}

func buildBoard() *board.Board {
	sps := buildShips()
	return board.New(sps)
}

func buildShips() []*ship.Ship {
	return []*ship.Ship{
		buildShip("A1", "A2", "A3"),
		buildShip("B1", "B2"),
	}
}

func buildShip(coordinates ...string) *ship.Ship {
	return ship.New(buildPositions(coordinates...))
}

func buildPositions(coordinates ...string) []*position.Position {
	var p []*position.Position
	for _, c := range coordinates {
		p = append(p, position.New(c))
	}
	return p
}
