package engine

import (
	"battleship/pkg/board"
	"battleship/pkg/game"
)

var instance *BattleShipGameEngine

type BattleShipGameEngine struct {
	b       board.Board
	games   []*game.Game
	winners []string
}

func New(b board.Board) *BattleShipGameEngine {
	if instance == nil {
		instance = &BattleShipGameEngine{
			b:       b,
			games:   []*game.Game{},
			winners: []string{},
		}
	}
	return instance
}

func (b *BattleShipGameEngine) Shoot(player, coordinates string) bool {
	g := b.getGameFor(player)
	isHit, finished := g.Shoot(coordinates)
	if finished {
		b.winners = append(b.winners, player)
	}
	return isHit
}

func (b *BattleShipGameEngine) TopTen() {
	panic("implement me!")
}

func (b *BattleShipGameEngine) getGameFor(player string) *game.Game {
	var g *game.Game
	for _, cg := range b.games {
		if cg.Player == player {
			g = cg
		}
	}
	if g == nil {
		nb := b.b
		g = game.New(player, &nb)
		b.games = append(b.games, g)
	}
	return g
}
