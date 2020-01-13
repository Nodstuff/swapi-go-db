package empire

import (
	"database/sql"
	"gopkg.in/guregu/null.v3"
)

type Vehicle struct {
	Id                   int         `json:"id,omitempty"`
	Name                 null.String `json:"name,omitempty"`
	Model                null.String `json:"model,omitempty"`
	Manufacturer         null.String `json:"manufacturer,omitempty"`
	CostInCredits        null.String `json:"cost_in_credits,omitempty"`
	Length               null.String `json:"length,omitempty"`
	MaxAtmospheringSpeed null.String `json:"max_atmosphering_speed,omitempty"`
	Crew                 null.String `json:"crew,omitempty"`
	Passengers           null.String `json:"passengers,omitempty"`
	CargoCapacity        null.String `json:"cargo_capacity,omitempty"`
	Consumables          null.String `json:"consumables,omitempty"`
	VehicleClass         null.String `json:"vehicle_class,omitempty"`
	Pilots               []Person    `json:"pilots,omitempty"`
	Films                []Film      `json:"films,omitempty"`
	Created              null.String `json:"created,omitempty"`
	Edited               null.String `json:"edited,omitempty"`
}

func (v *Vehicle) getPilots(db *sql.DB) {
	var pilots []Person
	rows, err := db.Query("select p.* from person p inner join vehicle_pilot sp on p.id = sp.person_id where vehicle_id = ?", v.Id)
	checkErr(err)

	defer rows.Close()

	for rows.Next() {
		var pilot Person
		rows.Scan(
			&pilot.Id,
			&pilot.Name,
			&pilot.Height,
			&pilot.Mass,
			&pilot.HairColor,
			&pilot.SkinColor,
			&pilot.EyeColor,
			&pilot.BirthYear,
			&pilot.Gender,
			&pilot.Homeworld,
			&pilot.Created,
			&pilot.Edited)
		pilots = append(pilots, pilot)
	}

	v.Pilots = pilots
}

func (v *Vehicle) getFilms(db *sql.DB) {
	var films []Film

	rows, err := db.Query("select f.* from film f inner join film_vehicle fc on f.id = fc.film_id where fc.vehicle_id = ?", v.Id)
	checkErr(err)
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

		checkErr(err)

		films = append(films, film)
	}

	v.Films = films
}
