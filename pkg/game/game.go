package game

import (
	"battleship/pkg/board"
)

type Game struct {
	Board    *board.Board
	Shots    int
	Player   string //TODO remove this one
	Finished bool
}

func New(player string, b *board.Board) *Game {
	return &Game{
		Board:    b,
		Shots:    0,
		Player:   player,
		Finished: false,
	}
}

func (g *Game) Shoot(coordinates string) (bool, bool) {
	g.Shots++
	isHit, allShipsSunk := g.Board.Shoot(coordinates)
	if allShipsSunk {
		g.Finished = true
	}
	return isHit, g.Finished
}
