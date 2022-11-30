package engine

import (
	"battleship/pkg/board"
	"battleship/pkg/game"
	"fmt"
	"sort"
	"sync"
)

var instance *BattleShipGameEngine

type BattleShipGameEngine struct {
	boardTemplate *board.Board //TODO tego pola nie powinno być, zamiast tego powinny być coordinates
	games         map[string]*game.Game
	gamesMutex    *sync.Mutex
	winners       []Winner
	winMutex      *sync.RWMutex
}

type Winner struct {
	Name  string
	Shots int
}

func New(b *board.Board) *BattleShipGameEngine {
	// TODO do konstrucktora powinienem przekazywać coordinates
	if instance == nil {
		instance = &BattleShipGameEngine{
			boardTemplate: b,
			games:         map[string]*game.Game{},
			gamesMutex:    &sync.Mutex{},
			winners:       []Winner{},
			winMutex:      &sync.RWMutex{},
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
		nb := board.Copy(ge.boardTemplate) //TODO tutaj powinienem zrobić nową instancję Board na podstawie koordynatów
		fmt.Printf("nb: %v\n", nb)
		fmt.Printf("ge.boardTemplate: %v\n", &ge.boardTemplate)
		ge.games[player] = game.New(player, nb)
	}
	return ge.games[player]
}
