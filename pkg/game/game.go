package game

import (
	"battleship/pkg/board"
)

type Game struct {
	B        *board.Board //TODO think of a better naming
	Player   string
	finished bool
}

func New(player string, b *board.Board) *Game {
	return &Game{
		B:        b,
		Player:   player,
		finished: false,
	}
}

// Shoot returns isHit and finished as a result of given shot
func (g *Game) Shoot(coordinates string) (bool, bool) {
	isHit, allShipsSunk := g.B.Shoot(coordinates)
	if allShipsSunk {
		g.finished = true
	}
	return isHit, g.finished
}
