package empire

import "fmt"

type Species struct {
	Name            string   `json:"name"`
	Classification  string   `json:"classification"`
	Designation     string   `json:"designation"`
	AverageHeight   string   `json:"average_height"`
	SkinColors      string   `json:"skin_colors"`
	HairColors      string   `json:"hair_colors"`
	EyeColors       string   `json:"eye_colors"`
	AverageLifespan string   `json:"average_lifespan"`
	Homeworld       string   `json:"homeworld"`
	Language        string   `json:"language"`
	People          []string `json:"people"`
	Films           []string `json:"films"`
	Created         string   `json:"created"`
	Edited          string   `json:"edited"`
	URL             string   `json:"url"`
}

func GetSpecies(id int) Species {
	var s Species
	GetHttp(fmt.Sprintf("/species/%d", id), &s)
	return s
}

func (s Species) GetPeople() []Person {
	return getPeople(s.People)
}

func (s Species) GetFilms() []Film {
	return getFilms(s.Films)
}
