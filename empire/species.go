package empire

import (
	"fmt"
	"gopkg.in/guregu/null.v3"
)

type Species struct {
	Id              int         `json:"id,omitempty"`
	Name            null.String `json:"name,omitempty"`
	Classification  null.String `json:"classification,omitempty"`
	Designation     null.String `json:"designation,omitempty"`
	AverageHeight   null.String `json:"average_height,omitempty"`
	SkinColors      null.String `json:"skin_colors,omitempty"`
	HairColors      null.String `json:"hair_colors,omitempty"`
	EyeColors       null.String `json:"eye_colors,omitempty"`
	AverageLifespan null.String `json:"average_lifespan,omitempty"`
	Homeworld       null.String `json:"homeworld,omitempty"`
	Language        null.String `json:"language,omitempty"`
	People          []Person    `json:"people,omitempty"`
	Films           []Film      `json:"films,omitempty"`
	Created         null.String `json:"created,omitempty"`
	Edited          null.String `json:"edited,omitempty"`
	URL             null.String `json:"url,omitempty"`
}

func GetSpecies(id int) Species {
	var s Species
	GetHttp(fmt.Sprintf("/species/%d", id), &s)
	return s
}
