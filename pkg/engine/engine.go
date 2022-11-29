package engine

import (
	"battleship/pkg/board"
	"battleship/pkg/game"
)

var instance *BattleShipGameEngine

type BattleShipGameEngine struct {
	b       *board.Board
	games   []*game.Game
	winners []string
}

func New(b *board.Board) *BattleShipGameEngine {
	if instance == nil {
		instance = &BattleShipGameEngine{
			b:       b,
			games:   []*game.Game{},
			winners: []string{},
		}
	}
	return instance
}

func (b *BattleShipGameEngine) Shoot(player, position string) (bool, error) {
	panic("implement me!")
}

func (b *BattleShipGameEngine) TopTen() {
	panic("implement me!")
}
