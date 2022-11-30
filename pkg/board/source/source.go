package source

import (
	"bufio"
	"os"
	"strings"
)

type CoordinatesSource interface {
	ReadCoordinates() ([][]string, error)
}

type CoordinatesFromFile struct {
	filename string
}

func New(filename string) *CoordinatesFromFile {
	return &CoordinatesFromFile{filename}
}

func (bs *CoordinatesFromFile) ReadCoordinates() ([][]string, error) {
	f, err := os.Open(bs.filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var shipsCoordinates [][]string
	for s.Scan() {
		line := s.Text()
		coordinates := strings.Split(line, ",")
		shipsCoordinates = append(shipsCoordinates, coordinates)
	}
	return shipsCoordinates, nil
}
