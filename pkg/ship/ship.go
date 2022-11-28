package ship

import (
	"battleship/pkg/position"
)

type Ship struct {
	positions []*position.Position
	sunk      bool
}

func New(positions []*position.Position) *Ship {
	return &Ship{positions, false}
}

func (s *Ship) Shoot(cooardinates string) (bool, bool) {
	var wasHit bool
	for _, p := range s.positions {
		if p.Coordinates == cooardinates {
			p.Shoot()
			wasHit = true
		}
	}
	if allPositionsAreHit(s) {
		s.sunk = true
	}
	return wasHit, s.sunk
}

func allPositionsAreHit(s *Ship) bool {
	for _, p := range s.positions {
		if !p.WasShot {
			return false
		}
	}
	return true
}
