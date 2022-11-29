package ship

import (
	"battleship/pkg/position"
)

type Ship struct {
	positions []*position.Position
	Sunk      bool
}

func New(positions []*position.Position) *Ship {
	// TODO change this constructor to accept coordinates ...string
	return &Ship{positions, false}
}

func (s *Ship) Shoot(cooardinates string) (bool, bool) {
	// TODO Ship nie powinien mieć funkcji Shoot - co powinien mieć w zamian?
	var wasHit bool
	for _, p := range s.positions {
		if p.Coordinates == cooardinates {
			p.Shoot()
			wasHit = true
		}
	}
	if allPositionsAreHit(s) {
		s.Sunk = true
	}
	return wasHit, s.Sunk
}

func allPositionsAreHit(s *Ship) bool {
	for _, p := range s.positions {
		if !p.WasShot {
			return false
		}
	}
	return true
}
