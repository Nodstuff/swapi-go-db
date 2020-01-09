package empire

import (
	"fmt"
	"gopkg.in/guregu/null.v3"
)

type Person struct {
	Id        int         `json:"id,omitempty"`
	Name      null.String `json:"name,omitempty"`
	Height    null.String `json:"height,omitempty"`
	Mass      null.String `json:"mass,omitempty"`
	HairColor null.String `json:"hair_color,omitempty"`
	SkinColor null.String `json:"skin_color,omitempty"`
	EyeColor  null.String `json:"eye_color,omitempty"`
	BirthYear null.String `json:"birth_year,omitempty"`
	Gender    null.String `json:"gender,omitempty"`
	Homeworld null.String `json:"homeworld,omitempty"`
	Films     []Film      `json:"films,omitempty"`
	Species   []Species   `json:"species,omitempty"`
	Vehicles  []Vehicle   `json:"vehicles,omitempty"`
	Starships []Starship  `json:"starships,omitempty"`
	Created   null.String `json:"created,omitempty"`
	Edited    null.String `json:"edited,omitempty"`
}

func GetPerson(id int) Person {
	var p Person
	GetHttp(fmt.Sprintf("/people/%d", id), &p)
	return p
}
