package ship

type Ship struct {
	Positions map[string]bool
	Sunk      bool
}

func New(coordinates ...string) *Ship {
	positions := make(map[string]bool)
	for _, c := range coordinates {
		positions[c] = false
	}
	return &Ship{positions, false}
}

func (s *Ship) MarkHit(cooardinates string) bool {
	s.Positions[cooardinates] = true
	if areAllPositionsHit(s) {
		s.Sunk = true
	}
	return s.Sunk
}

func areAllPositionsHit(s *Ship) bool {
	for _, p := range s.Positions {
		if !p {
			return false
		}
	}
	return true
}
