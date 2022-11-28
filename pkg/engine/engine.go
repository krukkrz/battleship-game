package engine

type BattleShipGameEngine struct {
	games   interface{}
	winners interface{}
}

func (b *BattleShipGameEngine) Shoot(player, position string) (bool, error) {
	panic("implement me!")
}

func (b *BattleShipGameEngine) TopTen() {
	panic("implement me!")
}
