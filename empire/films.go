package empire

import (
	"bufio"
	"database/sql"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"strings"
	"sync"
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

func (f *Film) getPeople(conn *sql.DB, wg *sync.WaitGroup) {
	go func() {
		var people []Person
		defer conn.Close()

		rows, err := conn.Query("select * from person where id in (select person_id from film_character where film_id = ?)", f.Id)
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
		f.Characters = people

		wg.Done()
	}()
}

func (f *Film) getPlanets(conn *sql.DB, wg *sync.WaitGroup) {
	go func() {
		var planets []Planet
		defer conn.Close()

		rows, err := conn.Query("select * from planet where id in (select planet_id from film_planet where film_id = ?)", f.Id)
		CheckErr(err)

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
		wg.Done()
	}()
}

func (f *Film) getStarships(conn *sql.DB, wg *sync.WaitGroup) {
	go func() {
		var ships []Starship
		defer conn.Close()

		rows, err := conn.Query("select * from starship where id in (select starship_id from film_starship where film_id = ?)", f.Id)
		CheckErr(err)

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

		wg.Done()
	}()
}

func (f *Film) getVehicles(conn *sql.DB, wg *sync.WaitGroup) {
	go func() {
		var vehicles []Vehicle
		defer conn.Close()

		rows, err := conn.Query("select * from vehicle where id in (select vehicle_id from film_vehicle where film_id = ?)", f.Id)
		CheckErr(err)

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

		wg.Done()
	}()
}

func (f *Film) getSpecies(conn *sql.DB, wg *sync.WaitGroup) {
	go func() {
		var species []Species
		defer conn.Close()

		rows, err := conn.Query("select * from species where id in (select species_id from film_species where film_id = ?)", f.Id)
		CheckErr(err)

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

		wg.Done()
	}()
}
