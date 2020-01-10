package empire

import (
	"database/sql"
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

func (p *Person) getFilms(db *sql.DB) {
	var films []Film

	rows, err := db.Query("select f.* from film f inner join film_character fc on f.id = fc.film_id where fc.person_id = ?", p.Id)
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

func (p *Person) getStarships(db *sql.DB) {
	var ships []Starship

	rows, err := db.Query("select s.* from starship s inner join starship_pilot sp on s.id = sp.starship_id where person_id = ?", p.Id)
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var ship Starship
		rows.Scan(
			&ship.Id,
			&ship.Name,
			&ship.Model,
			&ship.Manufacturer,
			&ship.CostInCredits,
			&ship.Length,
			&ship.MaxAtmospheringSpeed,
			&ship.Crew,
			&ship.Passengers,
			&ship.CargoCapacity,
			&ship.Consumables,
			&ship.HyperdriveRating,
			&ship.MGLT,
			&ship.StarshipClass,
			&ship.Created,
			&ship.Edited)
		ships = append(ships, ship)
	}

	p.Starships = ships
}

func (p *Person) getVehicles(db *sql.DB) {
	var vehicles []Vehicle

	rows, err := db.Query("select v.* from vehicle v inner join vehicle_pilot vp on v.id = vp.vehicle_id where person_id = ?", p.Id)
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var vehicle Vehicle
		rows.Scan(
			&vehicle.Id,
			&vehicle.Name,
			&vehicle.Model,
			&vehicle.Manufacturer,
			&vehicle.CostInCredits,
			&vehicle.Length,
			&vehicle.MaxAtmospheringSpeed,
			&vehicle.Crew,
			&vehicle.Passengers,
			&vehicle.CargoCapacity,
			&vehicle.Consumables,
			&vehicle.VehicleClass,
			&vehicle.Created,
			&vehicle.Edited)
		vehicles = append(vehicles, vehicle)
	}

	p.Vehicles = vehicles
}

func (p *Person) getSpecies(db *sql.DB) {
	var species []Species

	rows, err := db.Query("select s.* from species s inner join species_person sp on s.id = sp.species_id where person_id = ?", p.Id)
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var s Species
		rows.Scan(
			&s.Id,
			&s.Name,
			&s.Classification,
			&s.Designation,
			&s.AverageHeight,
			&s.SkinColors,
			&s.HairColors,
			&s.EyeColors,
			&s.AverageLifespan,
			&s.Homeworld,
			&s.Language,
			&s.Created,
			&s.Edited)
		species = append(species, s)
	}

	p.Species = species
}
