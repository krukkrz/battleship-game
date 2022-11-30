package board

import "battleship/pkg/ship"

type Board struct {
	Shots        int          //TODO move number of shots to game
	ships        []*ship.Ship //TODO change this to map[string]ShipType
	allShipsSunk bool
}

func New(ships []*ship.Ship) *Board {
	return &Board{
		Shots:        0,
		ships:        ships,
		allShipsSunk: false,
	}
}

func (b *Board) Shoot(cooardinates string) (bool, bool) {
	b.Shots++
	isHit := false
	for _, s := range b.ships {
		if _, ok := s.Positions[cooardinates]; ok {
			s.MarkHit(cooardinates)
			isHit = true
			break
		}
	}
	if allShipsSunk(b.ships) {
		b.allShipsSunk = true
	}
	return isHit, b.allShipsSunk
}

func allShipsSunk(ships []*ship.Ship) bool {
	for _, s := range ships {
		if !s.Sunk {
			return false
		}
	}
	return true
}
