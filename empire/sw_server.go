package empire

import (
	"encoding/json"
	"net/http"
)

const jsonContentType = "application/json"

type StarWarsServer struct {
	http.Handler
}

func NewStarWarsServer() *StarWarsServer {
	s := new(StarWarsServer)

	router := http.NewServeMux()
	router.Handle("/films", http.HandlerFunc(s.filmHandler))
	router.Handle("/people", http.HandlerFunc(s.peopleHandler))
	router.Handle("/planets", http.HandlerFunc(s.planetHandler))
	router.Handle("/species", http.HandlerFunc(s.speciesHandler))
	router.Handle("/starships", http.HandlerFunc(s.starshipHandler))
	router.Handle("/vehicles", http.HandlerFunc(s.vehicleHandler))

	s.Handler = router

	return s
}

func (s *StarWarsServer) filmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(GetFilms())
}

func (s *StarWarsServer) peopleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(GetPeople())
}

func (s *StarWarsServer) planetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(GetPlanets())
}

func (s *StarWarsServer) speciesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(GetSpecies())
}

func (s *StarWarsServer) starshipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(GetStarships())
}

func (s *StarWarsServer) vehicleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(GetVehicles())
}
