package empire

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

type Film struct {
	Id           int      `db:"-"`
	Title        string   `json:"title"`
	EpisodeId    int64    `json:"episode_id"`
	OpeningCrawl string   `json:"opening_crawl"`
	Director     string   `json:"director"`
	Producer     string   `json:"producer"`
	Characters   []string `json:"characters"`
	Planets      []string `json:"planets"`
	Starships    []string `json:"starships"`
	Vehicles     []string `json:"vehicles"`
	Species      []string `json:"species"`
	ReleaseDate  string   `json:"release_date"`
	Created      string   `json:"created"`
	Edited       string   `json:"edited"`
	URL          string   `json:"url"`
}

func GetFilm(id int) Film {
	var f Film
	GetHttp(fmt.Sprintf("/films/%d", id), &f)
	return f
}

func (f Film) PrintCrawl() {
	scanner := bufio.NewScanner(strings.NewReader(f.OpeningCrawl))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		time.Sleep(400 * time.Millisecond)
	}
}

func (f Film) GetCharacters() []Person {
	return getPeople(f.Characters)
}

func (f Film) GetPlanets() []Planet {
	return getPlanets(f.Planets)
}

func (f Film) GetStarships() []Starship {
	return getStarships(f.Starships)
}

func (f Film) GetVehicles() []Vehicle {
	return getVehicles(f.Vehicles)
}

func (f Film) GetSpecies() []Species {
	return getSpecies(f.Species)
}
