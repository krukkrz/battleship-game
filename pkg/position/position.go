package position

type Position struct {
	coordinates string
	wasShot     bool
}

func New(coordinates string) *Position {
	return &Position{coordinates, false}
}

func (p *Position) Shoot() bool {
	if !p.wasShot {
		p.wasShot = true
		return false
	}
	return p.wasShot
}
