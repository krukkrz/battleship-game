package engine

import (
	"battleship/pkg/board"
	"battleship/pkg/game"
	"sort"
	"sync"
)

var instance *BattleShipGameEngine

type BattleShipGameEngine struct {
	boardTemplate board.Board
	games         []*game.Game // TODO change this to map map[string]*game.Game
	gamesMutex    *sync.Mutex
	winners       []Winner
	winMutex      *sync.RWMutex
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
			gamesMutex:    &sync.Mutex{},
			winners:       []Winner{},
			winMutex:      &sync.RWMutex{},
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
	ge.winMutex.RLock()
	defer ge.winMutex.RUnlock()
	return ge.winners[:10]
}

func (ge *BattleShipGameEngine) addWinner(g *game.Game) {
	w := Winner{g.Player, g.Shots}
	ge.winMutex.Lock()
	defer ge.winMutex.Unlock()
	ge.winners = append(ge.winners, w)
	sort.SliceStable(ge.winners, func(i, j int) bool {
		return ge.winners[i].Shots > ge.winners[j].Shots
	})
}

func (ge *BattleShipGameEngine) getGameFor(player string) *game.Game {
	var g *game.Game
	ge.gamesMutex.Lock()
	defer ge.gamesMutex.Unlock()
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
