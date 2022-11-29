package position

type Position struct {
	Coordinates string
	WasShot     bool
}

func New(coordinates string) *Position {
	return &Position{coordinates, false}
}

// Shoot should be changed to sth like "mark"
func (p *Position) Shoot() bool {
	if !p.WasShot {
		p.WasShot = true
		return false
	}
	return p.WasShot
}
