package empire

import (
	"database/sql"
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

func (s *Species) getPeople(db *sql.DB) {
	var people []Person
	rows, err := db.Query("select p.* from person p inner join species_person sp on p.id = sp.person_id where species_id = ?", s.Id)
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

	s.People = people
}

func (s *Species) getFilms(db *sql.DB) {
	var films []Film

	rows, err := db.Query("select f.* from film f inner join film_species fs on f.id = fs.film_id where fs.species_id = ?", s.Id)
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

	s.Films = films
}
