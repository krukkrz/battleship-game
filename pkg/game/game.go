package game

import (
	"battleship/pkg/board"
)

type Game struct {
	Board    *board.Board
	Shots    int
	Finished bool
}

func New(b *board.Board) *Game {
	return &Game{
		Board:    b,
		Shots:    0,
		Finished: false,
	}
}

func (g *Game) Shoot(coordinates string) (bool, bool) {
	g.Shots++
	isHit, allShipsSunk := g.Board.Shoot(coordinates)
	if allShipsSunk {
		g.Finished = true
	}
	// TODO it should also return if a ship did sunk
	return isHit, g.Finished
}
