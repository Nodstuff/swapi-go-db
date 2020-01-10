package empire

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
)

const BaseURL = "https://swapi.co/api"

var ErrNotFound = errors.New("404: Not Found")

func GetHttp(path string, out interface{}) error {
	url := path

	if path[:4] != "http" {
		url = BaseURL + path
	}

	resp, err := http.Get(url)

	CheckErr(err)

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return ErrNotFound
	}

	body, err := ioutil.ReadAll(resp.Body)

	CheckErr(err)

	err = json.Unmarshal(body, &out)

	CheckErr(err)

	return nil
}

func getPeople(urls []string) (people []Person) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			var p Person
			GetHttp(url, &p)
			people = append(people, p)
			wg.Done()
		}(url)
	}
	wg.Wait()
	return
}

func getFilms(urls []string) (films []Film) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			var f Film
			GetHttp(url, &f)
			films = append(films, f)
			wg.Done()
		}(url)
	}
	wg.Wait()
	return
}

func getPlanets(urls []string) (planets []Planet) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			var p Planet
			GetHttp(url, &p)
			planets = append(planets, p)
			wg.Done()
		}(url)
	}
	wg.Wait()
	return
}

func getVehicles(urls []string) (vehicles []Vehicle) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			var v Vehicle
			GetHttp(url, &v)
			vehicles = append(vehicles, v)
			wg.Done()
		}(url)
	}
	wg.Wait()
	return
}

func getSpecies(urls []string) (species []Species) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			var s Species
			GetHttp(url, &s)
			species = append(species, s)
			wg.Done()
		}(url)
	}
	wg.Wait()
	return
}

func getStarships(urls []string) (starships []Starship) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			var s Starship
			GetHttp(url, &s)
			starships = append(starships, s)
			wg.Done()
		}(url)
	}
	wg.Wait()
	return
}

func PrettyPrintJson(data interface{}) []byte {
	prettyString, err := json.MarshalIndent(data, "", "    ")
	CheckErr(err)
	return prettyString
}
