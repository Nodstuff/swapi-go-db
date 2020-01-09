package empire

import (
	"fmt"
	"gopkg.in/guregu/null.v3"
)

type Planet struct {
	Id             int         `json:"id,omitempty"`
	Name           null.String `json:"name,omitempty"`
	RotationPeriod null.String `json:"rotation_period,omitempty"`
	OrbitalPeriod  null.String `json:"orbital_period,omitempty"`
	Diameter       null.String `json:"diameter,omitempty"`
	Climate        null.String `json:"climate,omitempty"`
	Gravity        null.String `json:"gravity,omitempty"`
	Terrain        null.String `json:"terrain,omitempty"`
	SurfaceWater   null.String `json:"surface_water,omitempty"`
	Population     null.String `json:"population,omitempty"`
	Residents      []Person    `json:"residents,omitempty"`
	Films          []Film      `json:"films,omitempty"`
	Created        null.String `json:"created,omitempty"`
	Edited         null.String `json:"edited,omitempty"`
}

func GetPlanet(id int) Planet {
	var p Planet
	GetHttp(fmt.Sprintf("/planets/%d", id), &p)
	return p
}

func (p Planet) getResidents() (people []Person) {
	conn := ConnectSql()
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM person WHERE homeworld = ?", p.Id)
	CheckErr(err)

	for rows.Next() {
		var person Person
		rows.Scan(
			&person.Id,
			&person.Name,
			&person.Height,
			&person.Mass,
			&person.HairColor,
			&person.SkinColor,
			&person.EyeColor,
			&person.BirthYear,
			&person.Gender,
			&person.Homeworld,
			&person.Created,
			&person.Edited)
		people = append(people, person)
	}
	return
}
