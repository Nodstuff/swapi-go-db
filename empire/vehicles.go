package empire

import (
	"fmt"
)

type Vehicle struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	VehicleClass         string   `json:"vehicle_class"`
	Pilots               []string `json:"pilots"`
	Films                []string `json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	URL                  string   `json:"url"`
}

func GetVehicle(id int) Vehicle {
	var v Vehicle
	GetHttp(fmt.Sprintf("/vehicles/%d", id), &v)
	return v
}

func (v Vehicle) GetPilots() []Person {
	return getPeople(v.Pilots)
}

func (v Vehicle) GetFilms() []Film {
	return getFilms(v.Films)
}
