package main

import (
	"battleship/pkg/board/source"
	"battleship/pkg/engine"
	"battleship/pkg/server"
	"log"
)

func main() {
	// TODO start a web serwer and run games
	// TODO read positions from txt file

	c, err := source.New("positions.txt").ReadCoordinates()
	if err != nil {
		log.Fatalf("cannot read coordinates, error: %v", err)
	}
	e := engine.New(c)
	s := server.New(":8080", e)
	s.RegisterRoutes()
	s.Run()
}
