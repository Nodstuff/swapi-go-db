package empire

import "fmt"

type Person struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	Homeworld string   `json:"homeworld"`
	Films     []string `json:"films"`
	Species   []string `json:"species"`
	Vehicles  []string `json:"vehicles"`
	Starships []string `json:"starships"`
	Created   string   `json:"created"`
	Edited    string   `json:"edited"`
	URL       string   `json:"url"`
}

func GetPerson(id int) Person {
	var p Person
	GetHttp(fmt.Sprintf("/people/%d", id), &p)
	return p
}

func (p Person) GetFilms() []string {
	return p.Films
}

func (p Person) GetSpecies() []Species {
	return getSpecies(p.Species)
}

func (p Person) GetVehicles() []Vehicle {
	return getVehicles(p.Vehicles)
}

func (p Person) GetStarships() []Starship {
	return getStarships(p.Starships)
}
