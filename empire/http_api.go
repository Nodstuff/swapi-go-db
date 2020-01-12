package empire

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
)

const baseURL = "https://swapi.co/api"

var errNotFound = errors.New("404: Not Found")

func getHttp(path string, out interface{}) error {
	url := path

	if path[:4] != "http" {
		url = baseURL + path
	}

	resp, err := http.Get(url)

	checkErr(err)

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return errNotFound
	}

	body, err := ioutil.ReadAll(resp.Body)

	checkErr(err)

	err = json.Unmarshal(body, &out)

	checkErr(err)

	return nil
}

func getPeopleHttp(urls []string) (people []Person) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			var p Person
			getHttp(url, &p)
			people = append(people, p)
			wg.Done()
		}(url)
	}
	wg.Wait()
	return
}

func getFilmsHttp(urls []string) (films []Film) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			var f Film
			getHttp(url, &f)
			films = append(films, f)
			wg.Done()
		}(url)
	}
	wg.Wait()
	return
}

func getPlanetsHttp(urls []string) (planets []Planet) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			var p Planet
			getHttp(url, &p)
			planets = append(planets, p)
			wg.Done()
		}(url)
	}
	wg.Wait()
	return
}

func getVehiclesHttp(urls []string) (vehicles []Vehicle) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			var v Vehicle
			getHttp(url, &v)
			vehicles = append(vehicles, v)
			wg.Done()
		}(url)
	}
	wg.Wait()
	return
}

func getSpeciesHttp(urls []string) (species []Species) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			var s Species
			getHttp(url, &s)
			species = append(species, s)
			wg.Done()
		}(url)
	}
	wg.Wait()
	return
}

func getStarshipsHttp(urls []string) (starships []Starship) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			var s Starship
			getHttp(url, &s)
			starships = append(starships, s)
			wg.Done()
		}(url)
	}
	wg.Wait()
	return
}

func prettyPrintJson(data interface{}) []byte {
	prettyString, err := json.MarshalIndent(data, "", "    ")
	checkErr(err)
	return prettyString
}
