package engine

import (
	"battleship/pkg/board"
	"battleship/pkg/game"
	"sort"
)

var instance *BattleShipGameEngine

type BattleShipGameEngine struct {
	b       board.Board //TODO think of a better naming
	games   []*game.Game
	winners []Winner
}

type Winner struct {
	Name  string
	Shots int
}

func New(b board.Board) *BattleShipGameEngine {
	if instance == nil {
		instance = &BattleShipGameEngine{
			b:       b,
			games:   []*game.Game{},
			winners: []Winner{},
		}
	}
	return instance
}

func (b *BattleShipGameEngine) Shoot(player, coordinates string) bool {
	g := b.getGameFor(player)
	isHit, finished := g.Shoot(coordinates)
	if finished {
		b.addWinner(g)
	}
	return isHit
}

func (b *BattleShipGameEngine) TopTen() []Winner {
	return b.winners[:10]
}

func (b *BattleShipGameEngine) addWinner(g *game.Game) {
	w := Winner{g.Player, g.B.Shots}
	b.winners = append(b.winners, w)
	sort.SliceStable(b.winners, func(i, j int) bool {
		return b.winners[i].Shots > b.winners[j].Shots
	})
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
