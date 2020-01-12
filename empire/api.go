package empire

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"swapi-go-db/empire/database"
)

func ConnectDB() {
	var err error
	database.DBCon, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/sw_data")
	CheckErr(err)
}

func GetFilms() (films []Film) {
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

	return
}

func GetPeople() (people []Person) {
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

	return
}

func GetPlanets() (planets []Planet) {
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

	return
}

func GetSpecies() (species []Species) {
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

	return
}

func GetStarships() (ships []Starship) {
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

	return
}

func GetVehicles() (vehicles []Vehicle) {
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

	return
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
