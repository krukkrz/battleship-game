package board

import "battleship/pkg/ship"

type Board struct {
	ships        []*ship.Ship
	allShipsSunk bool
}

func New(ships []*ship.Ship) *Board {
	return &Board{
		ships:        ships,
		allShipsSunk: false,
	}
}

func (b *Board) Shoot(cooardinates string) (bool, bool) {
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
