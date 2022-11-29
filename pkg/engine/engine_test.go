package engine

import (
	"battleship/pkg/board"
	"battleship/pkg/game"
	"battleship/pkg/position"
	"battleship/pkg/ship"
	"testing"
)

func TestBattleShipGameEngine_New(t *testing.T) {
	b := buildBoard()
	e1 := New(b)
	e2 := New(b)
	if e1 != e2 {
		t.Error("BattleShipGameEngine should be singleton but is not")
	}
}

// do concurrent tests of using BattleShipGameEngine

func TestBattleShipGameEngine_Shoot_creates_new_game_for_given_player(t *testing.T) {
	b := buildBoard()
	e := New(b)

	e.Shoot("mark", "A1")

	if len(e.games) != 1 {
		t.Errorf("should have one game for mark, but had %d games", len(e.games))
	}

	cleanup()
}

func TestBattleShipGameEngine_Shoot_adds_shot_for_given_player(t *testing.T) {
	b := buildBoard()
	e := New(b)

	e.Shoot("mark", "A1")
	e.Shoot("edi", "A1")
	e.Shoot("edi", "A2")

	if len(e.games) != 2 {
		t.Errorf("should have two games, but had %d games", len(e.games))
	}

	ediGame := getGameFor(e.games, "edi")
	if ediGame.B.Shots != 2 {
		t.Errorf("expected two shots for edi, got: %d", ediGame.B.Shots)
	}

	markGame := getGameFor(e.games, "mark")
	if markGame.B.Shots != 1 {
		t.Errorf("expected one shot for mark, got: %d", markGame.B.Shots)
	}

	cleanup()
}

func TestBattleShipGameEngine_Shoot_adds_players_to_winners_when_finished(t *testing.T) {
	b := buildBoard()
	e := New(b)

	e.Shoot("mark", "A1")
	e.Shoot("edi", "A1")
	e.Shoot("edi", "A2")
	e.Shoot("edi", "A3")
	e.Shoot("edi", "B1")
	e.Shoot("edi", "B2")

	if len(e.winners) != 1 {
		t.Errorf("expected one winner, got: %d", len(e.winners))
	}

	cleanup()
}

func TestBattleShipGameEngine_TopTen(t *testing.T) {
	// returns top ten winners
}

// move those builders to common/test
func buildBoard() board.Board {
	sps := buildShips()
	return *board.New(sps)
}

func buildShips() []*ship.Ship {
	return []*ship.Ship{
		buildShip("A1", "A2", "A3"),
		buildShip("B1", "B2"),
	}
}

func buildShip(coordinates ...string) *ship.Ship {
	return ship.New(buildPositions(coordinates...))
}

func buildPositions(coordinates ...string) []*position.Position {
	var p []*position.Position
	for _, c := range coordinates {
		p = append(p, position.New(c))
	}
	return p
}

func getGameFor(games []*game.Game, player string) *game.Game {
	var g *game.Game
	for _, cg := range games {
		if cg.Player == player {
			g = cg
		}
	}
	return g
}

func cleanup() {
	instance = nil
}
