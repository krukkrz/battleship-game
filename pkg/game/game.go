package game

import (
	"battleship/pkg/board"
)

type Game struct {
	Board    *board.Board
	Player   string
	finished bool
}

func New(player string, b *board.Board) *Game {
	return &Game{
		Board:    b,
		Player:   player,
		finished: false,
	}
}

func (g *Game) Shoot(coordinates string) (bool, bool) {
	isHit, allShipsSunk := g.Board.Shoot(coordinates)
	if allShipsSunk {
		g.finished = true
	}
	return isHit, g.finished
}
