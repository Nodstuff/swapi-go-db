package empire

import (
	"bufio"
	"database/sql"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"strings"
	"time"
)

type Film struct {
	Id           int         `db:"-,omitempty"`
	Title        null.String `json:"title,omitempty"`
	EpisodeId    int64       `json:"episode_id,omitempty"`
	OpeningCrawl null.String `json:"opening_crawl,omitempty"`
	Director     null.String `json:"director,omitempty"`
	Producer     null.String `json:"producer,omitempty"`
	Characters   []Person    `json:"characters,omitempty"`
	Planets      []Planet    `json:"planets,omitempty"`
	Starships    []Starship  `json:"starships,omitempty"`
	Vehicles     []Vehicle   `json:"vehicles,omitempty"`
	Species      []Species   `json:"species,omitempty"`
	ReleaseDate  null.String `json:"release_date,omitempty"`
	Created      null.String `json:"created,omitempty"`
	Edited       null.String `json:"edited,omitempty"`
}

func (f Film) PrintCrawl() {
	scanner := bufio.NewScanner(strings.NewReader(f.OpeningCrawl.String))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		time.Sleep(400 * time.Millisecond)
	}
}

func (f *Film) getPeople(db *sql.DB) {
	var people []Person

	rows, err := db.Query("select p.* from Person p inner join film_character fc on p.id = fc.person_id and fc.film_id = ?", f.Id)
	checkErr(err)
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

	f.Characters = people
}

func (f *Film) getPlanets(db *sql.DB) {
	var planets []Planet

	rows, err := db.Query("select p.* from planet p inner join film_planet fp on p.id = fp.planet_id where film_id = ?", f.Id)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var planet Planet
		rows.Scan(
			&planet.Id,
			&planet.Name,
			&planet.RotationPeriod,
			&planet.OrbitalPeriod,
			&planet.Diameter,
			&planet.Climate,
			&planet.Gravity,
			&planet.Terrain,
			&planet.SurfaceWater,
			&planet.Population,
			&planet.Created,
			&planet.Edited)
		planets = append(planets, planet)
	}

	f.Planets = planets
}

func (f *Film) getStarships(db *sql.DB) {
	var ships []Starship

	rows, err := db.Query("select s.* from starship s inner join film_starship fs on s.id = fs.starship_id and fs.film_id = ?", f.Id)
	checkErr(err)
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

	f.Starships = ships
}

func (f *Film) getVehicles(db *sql.DB) {
	var vehicles []Vehicle

	rows, err := db.Query("select v.* from vehicle v inner join film_vehicle fv on v.id = fv.vehicle_id and fv.film_id = ?", f.Id)
	checkErr(err)
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

	f.Vehicles = vehicles
}

func (f *Film) getSpecies(db *sql.DB) {
	var species []Species

	rows, err := db.Query("select s.* from species s inner join film_species fs on s.id = fs.species_id and fs.film_id = ?", f.Id)
	checkErr(err)
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

	f.Species = species
}
