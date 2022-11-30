package engine

import (
	"battleship/pkg/common/test"
	"fmt"
	"sync"
	"testing"
)

func TestBattleShipGameEngine_Shoot_concurrent(t *testing.T) {
	e := New(test.Coordinates)
	n := 2
	wg := &sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		player := fmt.Sprintf("player-%d", i)
		go func(e *BattleShipGameEngine, player string) {
			defer wg.Done()
			e.Shoot(player, "A1")
			e.Shoot(player, "A2")
			e.Shoot(player, "A3")
			e.Shoot(player, "B1")
			e.Shoot(player, "B2")

			g := e.games[player]
			if g.Shots != 5 {
				t.Errorf("expected five shots for %s, got: %d", player, g.Shots)
			}
		}(e, player)
	}

	wg.Wait()

	if len(e.games) != n {
		t.Errorf("expected %d number of games, got: %d", n, len(e.games))
	}

	if len(e.winners) != n {
		for _, w := range e.winners {
			println(w.Name)
		}
		t.Errorf("expected %d number of winners, got: %d", n, len(e.winners))
	}
	cleanup()
}

func TestBattleShipGameEngine_New(t *testing.T) {
	e1 := New(test.Coordinates)
	e2 := New(test.Coordinates)
	if e1 != e2 {
		t.Error("BattleShipGameEngine should be singleton but is not")
	}
	cleanup()
}

func TestBattleShipGameEngine_Shoot_creates_new_game_for_given_player(t *testing.T) {
	e := New(test.Coordinates)

	e.Shoot("mark", "A1")

	if len(e.games) != 1 {
		t.Errorf("should have one game for mark, but had %d games", len(e.games))
	}

	cleanup()
}

func TestBattleShipGameEngine_Shoot_adds_shot_for_given_player(t *testing.T) {
	e := New(test.Coordinates)

	e.Shoot("mark", "A1")
	e.Shoot("edi", "A1")
	e.Shoot("edi", "A2")

	if len(e.games) != 2 {
		t.Errorf("should have two games, but had %d games", len(e.games))
	}

	ediGame := e.games["edi"]
	if ediGame.Shots != 2 {
		t.Errorf("expected two shots for edi, got: %d", ediGame.Shots)
	}

	markGame := e.games["mark"]
	if markGame.Shots != 1 {
		t.Errorf("expected one shot for mark, got: %d", markGame.Shots)
	}

	cleanup()
}

func TestBattleShipGameEngine_Shoot_adds_players_to_winners_when_finished(t *testing.T) {
	e := New(test.Coordinates)

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
	e := New(test.Coordinates)

	buildElevenPlayers(e)

	ws := e.TopTen()

	if len(ws) > 10 {
		t.Errorf("there should be only ten winners in top ten, got: %d", len(ws))
	}

	for i, w := range ws {
		if i == 0 {
			continue
		}
		if w.Shots > ws[i-1].Shots {
			t.Errorf("%s has more votes than %s but it should not - sorting is not done well", w.Name, ws[i-1].Name)
		}
	}

	cleanup()
}

func TestBattleShipGameEngine_Shoot_returns_errror_if_game_is_finished(t *testing.T) {
	e := New(test.Coordinates)

	edi := "edi"
	e.Shoot(edi, "A1")
	e.Shoot(edi, "A2")
	e.Shoot(edi, "A3")
	e.Shoot(edi, "B1")
	e.Shoot(edi, "B2")

	_, err := e.Shoot(edi, "C2")

	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}

func buildElevenPlayers(e *BattleShipGameEngine) {
	barry2 := "barry2"
	e.Shoot(barry2, "A1")
	e.Shoot(barry2, "A2")
	e.Shoot(barry2, "C2")
	e.Shoot(barry2, "C2")
	e.Shoot(barry2, "C2")
	e.Shoot(barry2, "C2")
	e.Shoot(barry2, "C2")
	e.Shoot(barry2, "C2")
	e.Shoot(barry2, "C2")
	e.Shoot(barry2, "C2")
	e.Shoot(barry2, "C2")
	e.Shoot(barry2, "A3")
	e.Shoot(barry2, "B1")
	e.Shoot(barry2, "B2")

	marry2 := "marry2"
	e.Shoot(marry2, "A1")
	e.Shoot(marry2, "A2")
	e.Shoot(marry2, "C2")
	e.Shoot(marry2, "C2")
	e.Shoot(marry2, "C2")
	e.Shoot(marry2, "C2")
	e.Shoot(marry2, "C2")
	e.Shoot(marry2, "C2")
	e.Shoot(marry2, "C2")
	e.Shoot(marry2, "C2")
	e.Shoot(marry2, "C2")
	e.Shoot(marry2, "A3")
	e.Shoot(marry2, "B1")
	e.Shoot(marry2, "B2")

	bob2 := "bob2"
	e.Shoot(bob2, "A1")
	e.Shoot(bob2, "A2")
	e.Shoot(bob2, "C2")
	e.Shoot(bob2, "C2")
	e.Shoot(bob2, "C2")
	e.Shoot(bob2, "C2")
	e.Shoot(bob2, "C2")
	e.Shoot(bob2, "C2")
	e.Shoot(bob2, "C2")
	e.Shoot(bob2, "C2")
	e.Shoot(bob2, "A3")
	e.Shoot(bob2, "B1")
	e.Shoot(bob2, "B2")

	mark2 := "mark2"
	e.Shoot(mark2, "A1")
	e.Shoot(mark2, "A2")
	e.Shoot(mark2, "C2")
	e.Shoot(mark2, "C2")
	e.Shoot(mark2, "C2")
	e.Shoot(mark2, "C2")
	e.Shoot(mark2, "C2")
	e.Shoot(mark2, "C2")
	e.Shoot(mark2, "A3")
	e.Shoot(mark2, "B1")
	e.Shoot(mark2, "B2")

	benny2 := "benny2"
	e.Shoot(benny2, "A1")
	e.Shoot(benny2, "A2")
	e.Shoot(benny2, "C2")
	e.Shoot(benny2, "C2")
	e.Shoot(benny2, "C2")
	e.Shoot(benny2, "C2")
	e.Shoot(benny2, "C2")
	e.Shoot(benny2, "C2")
	e.Shoot(benny2, "C2")
	e.Shoot(benny2, "A3")
	e.Shoot(benny2, "B1")
	e.Shoot(benny2, "B2")

	edi2 := "edi2"
	e.Shoot(edi2, "A1")
	e.Shoot(edi2, "A2")
	e.Shoot(edi2, "C2")
	e.Shoot(edi2, "C2")
	e.Shoot(edi2, "C2")
	e.Shoot(edi2, "C2")
	e.Shoot(edi2, "C2")
	e.Shoot(edi2, "A3")
	e.Shoot(edi2, "B1")
	e.Shoot(edi2, "B2")

	bob := "bob"
	e.Shoot(bob, "A1")
	e.Shoot(bob, "A2")
	e.Shoot(bob, "C2")
	e.Shoot(bob, "C2")
	e.Shoot(bob, "C2")
	e.Shoot(bob, "A3")
	e.Shoot(bob, "B1")
	e.Shoot(bob, "B2")

	marry := "marry"
	e.Shoot(marry, "A1")
	e.Shoot(marry, "A2")
	e.Shoot(marry, "C2")
	e.Shoot(marry, "C2")
	e.Shoot(marry, "C2")
	e.Shoot(marry, "C2")
	e.Shoot(marry, "A3")
	e.Shoot(marry, "B1")
	e.Shoot(marry, "B2")

	benny := "benny"
	e.Shoot(benny, "A1")
	e.Shoot(benny, "A2")
	e.Shoot(benny, "C2")
	e.Shoot(benny, "C2")
	e.Shoot(benny, "A3")
	e.Shoot(benny, "B1")
	e.Shoot(benny, "B2")

	mark := "mark"
	e.Shoot(mark, "A1")
	e.Shoot(mark, "A2")
	e.Shoot(mark, "C2")
	e.Shoot(mark, "A3")
	e.Shoot(mark, "B1")
	e.Shoot(mark, "B2")

	edi := "edi"
	e.Shoot(edi, "A1")
	e.Shoot(edi, "A2")
	e.Shoot(edi, "A3")
	e.Shoot(edi, "B1")
	e.Shoot(edi, "B2")
}

func cleanup() {
	instance = nil
}
