package engine

import (
	"battleship/pkg/board"
	"battleship/pkg/game"
	"battleship/pkg/ship"
	"fmt"
	"sort"
	"sync"
)

var instance *BattleShipGameEngine

type BattleShipGameEngine struct {
	coordinates [][]string
	games       map[string]*game.Game
	gamesMutex  *sync.Mutex
	winners     []Winner
	winMutex    *sync.RWMutex
}

type Winner struct {
	Name  string
	Shots int
}

func New(coordinates [][]string) *BattleShipGameEngine {
	if instance == nil {
		instance = &BattleShipGameEngine{
			coordinates: coordinates,
			games:       map[string]*game.Game{},
			gamesMutex:  &sync.Mutex{},
			winners:     []Winner{},
			winMutex:    &sync.RWMutex{},
		}
	}
	return instance
}

func (ge *BattleShipGameEngine) Shoot(player, coordinates string) (bool, error) {
	g := ge.getGameFor(player)
	if g.Finished {
		return false, fmt.Errorf("game is already finished")
	}
	isHit, finished := g.Shoot(coordinates)
	if finished {
		ge.addWinner(player, g)
	}
	return isHit, nil
}

func (ge *BattleShipGameEngine) TopTen() []Winner {
	ge.winMutex.RLock()
	defer ge.winMutex.RUnlock()
	return ge.winners[:10]
}

func (ge *BattleShipGameEngine) addWinner(player string, g *game.Game) {
	w := Winner{player, g.Shots}
	ge.winMutex.Lock()
	defer ge.winMutex.Unlock()
	ge.winners = append(ge.winners, w)
	sort.SliceStable(ge.winners, func(i, j int) bool {
		return ge.winners[i].Shots > ge.winners[j].Shots
	})
}

func (ge *BattleShipGameEngine) getGameFor(player string) *game.Game {
	ge.gamesMutex.Lock()
	defer ge.gamesMutex.Unlock()
	if _, ok := ge.games[player]; !ok {

		var ships []*ship.Ship
		for _, c := range ge.coordinates {
			ships = append(ships, ship.New(c...))
		}

		nb := board.New(ships)
		ge.games[player] = game.New(nb)
	}
	return ge.games[player]
}
