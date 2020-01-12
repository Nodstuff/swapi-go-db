package empire

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"swapi-go-db/empire/database"
)

var cacheMap = make(map[string][]interface{})

func ConnectDB() {
	var err error
	database.DBCon, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/sw_data")
	CheckErr(err)
}

func GetFilms() (films []interface{}) {

	if res, ok := cacheMap["films"]; ok {
		return res
	}

	db := database.DBCon

	rows, err := db.Query("SELECT * FROM film")
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

		film.getPeople(db)
		film.getPlanets(db)
		film.getStarships(db)
		film.getVehicles(db)
		film.getSpecies(db)

		films = append(films, film)
	}

	cacheMap["films"] = films

	return
}

func GetPeople() (people []interface{}) {

	if res, ok := cacheMap["people"]; ok {
		return res
	}

	db := database.DBCon

	rows, err := db.Query("SELECT * FROM person")
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var person Person
		err := rows.Scan(
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

		CheckErr(err)

		person.getStarships(db)
		person.getVehicles(db)
		person.getSpecies(db)

		people = append(people, person)
	}

	cacheMap["people"] = people

	return
}

func GetPlanets() (planets []interface{}) {

	if res, ok := cacheMap["planets"]; ok {
		return res
	}

	db := database.DBCon

	rows, err := db.Query("SELECT * FROM planet")
	CheckErr(err)
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

		planet.getFilms(db)
		planet.getResidents(db)

		planets = append(planets, planet)
	}

	cacheMap["planets"] = planets

	return
}

func GetSpecies() (species []interface{}) {

	if res, ok := cacheMap["species"]; ok {
		return res
	}

	db := database.DBCon

	rows, err := db.Query("SELECT * FROM species")
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

		s.getFilms(db)
		s.getPeople(db)

		species = append(species, s)
	}

	cacheMap["species"] = species

	return
}

func GetStarships() (ships []interface{}) {

	if res, ok := cacheMap["starships"]; ok {
		return res
	}

	db := database.DBCon

	rows, err := db.Query("SELECT * FROM starship")
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

		ship.getFilms(db)
		ship.getPilots(db)

		ships = append(ships, ship)
	}

	cacheMap["starships"] = ships

	return
}

func GetVehicles() (vehicles []interface{}) {

	if res, ok := cacheMap["vehicles"]; ok {
		return res
	}

	db := database.DBCon

	rows, err := db.Query("SELECT * FROM vehicle")
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

		vehicle.getFilms(db)
		vehicle.getPilots(db)

		vehicles = append(vehicles, vehicle)
	}

	cacheMap["vehicles"] = vehicles

	return
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
