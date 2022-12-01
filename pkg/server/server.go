package server

import (
	"battleship/pkg/engine"
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	port string
	e    *engine.BattleShipGameEngine
	mux  *http.ServeMux
}

func New(port string, e *engine.BattleShipGameEngine) *Server {
	m := http.NewServeMux()
	return &Server{port, e, m}
}

func (s *Server) Run() {
	log.Printf("started server on port %s \n", s.port)
	http.ListenAndServe(s.port, s.mux)
}

func (s *Server) RegisterRoutes() *http.ServeMux {
	log.Println("registering routes...")
	s.mux.HandleFunc("/shoot", s.HandleShoot)
	s.mux.HandleFunc("/top-ten", s.HandleTopTen)
	return s.mux
}

func (s *Server) HandleTopTen(w http.ResponseWriter, r *http.Request) {
	type responseItem struct {
		Name  string `json:"name"`
		Shots int    `json:"shots"`
	}

	winners := s.e.TopTen()

	var winnersResponse []responseItem
	for _, winner := range winners {
		winnersResponse = append(winnersResponse, responseItem{winner.Name, winner.Shots})
	}

	resp, err := json.Marshal(winnersResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp))
}

func (s *Server) HandleShoot(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Player      string `json:"player"`
		Coordinates string `json:"coordinates"`
	}

	type response struct {
		IsHit        bool `json:"isHit"`
		GameFinished bool `json:"gameFinished"`
	}

	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isHit, isFinished, err := s.e.Shoot(req.Player, req.Coordinates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(response{isHit, isFinished})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp))
}
