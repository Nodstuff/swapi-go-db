package empire

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

const jsonContentType = "application/json"

type StarWarsServer struct {
	http.Handler
}

func NewStarWarsServer() *StarWarsServer {
	s := new(StarWarsServer)

	router := mux.NewRouter()

	router.HandleFunc("/films", s.filmHandler)
	router.HandleFunc("/people", s.peopleHandler)
	router.HandleFunc("/planets", s.planetHandler)
	router.HandleFunc("/species", s.speciesHandler)
	router.HandleFunc("/starships", s.starshipHandler)
	router.HandleFunc("/vehicles", s.vehicleHandler)

	s.Handler = handlers.CompressHandler(router)

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
