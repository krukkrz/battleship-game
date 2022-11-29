package game

import (
	"battleship/pkg/board"
)

type Game struct {
	b        *board.Board
	player   string
	finished bool
}

func New(player string, b *board.Board) *Game {
	return &Game{
		b:        b,
		player:   player,
		finished: false,
	}
}

func (g *Game) Shoot(coordinates string) (bool, bool) {
	isHit, allShipsSunk := g.b.Shoot(coordinates)
	if allShipsSunk {
		g.finished = true
	}
	return isHit, g.finished
}
