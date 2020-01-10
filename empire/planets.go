package empire

import (
	"database/sql"
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

func (p *Planet) getResidents(db *sql.DB) {
	var people []Person
	rows, err := db.Query("SELECT * FROM person WHERE homeworld = ?", p.Id)
	CheckErr(err)

	defer rows.Close()

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

	p.Residents = people
}

func (p *Planet) getFilms(db *sql.DB) {
	var films []Film

	rows, err := db.Query("select f.* from film f inner join film_planet fc on f.id = fc.film_id where fc.planet_id = ?", p.Id)
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var film Film

		err := rows.Scan(
			&film.Id,
			&film.Title,
			&film.EpisodeId,
			&film.OpeningCrawl,
			&film.Director,
			&film.Producer,
			&film.ReleaseDate,
			&film.Created,
			&film.Edited)

		CheckErr(err)

		films = append(films, film)
	}

	p.Films = films
}
