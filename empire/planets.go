package empire

import (
	"fmt"
)

type Planet struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	Residents      []string `json:"residents"`
	Films          []string `json:"films"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	URL            string   `json:"url"`
}

func GetPlanet(id int) Planet {
	var p Planet
	GetHttp(fmt.Sprintf("/planets/%d", id), &p)
	return p
}

func (p Planet) GetResidents() []Person {
	return getPeople(p.Residents)
}

func (p Planet) GetFilms() []Film {
	return getFilms(p.Films)
}
