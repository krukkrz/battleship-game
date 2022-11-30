package source

import (
	"battleship/pkg/board"
	"battleship/pkg/ship"
	"bufio"
	"os"
	"strings"
)

type BoardSource interface {
	SetupBoard() (*board.Board, error)
}

type BoardFromFile struct {
	filename string
}

func New(filename string) *BoardFromFile {
	return &BoardFromFile{filename}
}

func (bs *BoardFromFile) SetupBoard() (*board.Board, error) {
	f, err := os.Open(bs.filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var ships []*ship.Ship
	for s.Scan() {
		line := s.Text()
		coordinates := strings.Split(line, ",")
		ships = append(ships, ship.New(coordinates...))
	}
	return board.New(ships), nil
}
