package engine

import (
	"battleship/pkg/board"
	"battleship/pkg/game"
	"sort"
)

var instance *BattleShipGameEngine

type BattleShipGameEngine struct {
	boardTemplate board.Board
	games         []*game.Game
	winners       []Winner
}

type Winner struct {
	Name  string
	Shots int
}

func New(b board.Board) *BattleShipGameEngine {
	if instance == nil {
		instance = &BattleShipGameEngine{
			boardTemplate: b,
			games:         []*game.Game{},
			winners:       []Winner{},
		}
	}
	return instance
}

func (ge *BattleShipGameEngine) Shoot(player, coordinates string) bool {
	g := ge.getGameFor(player)
	isHit, finished := g.Shoot(coordinates)
	if finished {
		ge.addWinner(g)
	}
	return isHit
}

func (ge *BattleShipGameEngine) TopTen() []Winner {
	return ge.winners[:10]
}

func (ge *BattleShipGameEngine) addWinner(g *game.Game) {
	w := Winner{g.Player, g.B.Shots}
	ge.winners = append(ge.winners, w)
	sort.SliceStable(ge.winners, func(i, j int) bool {
		return ge.winners[i].Shots > ge.winners[j].Shots
	})
}

func (ge *BattleShipGameEngine) getGameFor(player string) *game.Game {
	var g *game.Game
	for _, cg := range ge.games {
		if cg.Player == player {
			g = cg
		}
	}
	if g == nil {
		nb := ge.boardTemplate
		g = game.New(player, &nb)
		ge.games = append(ge.games, g)
	}
	return g
}
