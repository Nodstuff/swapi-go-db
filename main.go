package main

import (
	"fmt"
	"github.com/fatih/color"
	"helloWorld/empire"
	"sync"
)

func main() {
	empire.ConnectSql()
}

func getFilmAndPrintCrawl(id int, printColor color.Attribute) {
	color.Set(printColor)
	empire.GetFilm(id).PrintCrawl()
}

func getFilmAndNestedObjects(id int) {
	var wg sync.WaitGroup
	wg.Add(4)

	var vehicles []empire.Vehicle
	var starships []empire.Starship
	var characters []empire.Person
	var planets []empire.Planet

	film := empire.GetFilm(id)

	go func(film empire.Film) {
		vehicles = film.GetVehicles()
		wg.Done()
	}(film)

	go func(film empire.Film) {
		starships = film.GetStarships()
		wg.Done()
	}(film)

	go func(film empire.Film) {
		characters = film.GetCharacters()
		wg.Done()
	}(film)

	go func(film empire.Film) {
		planets = film.GetPlanets()
		wg.Done()
	}(film)

	wg.Wait()

	fmt.Println(empire.PrettyPrintJson(vehicles))
	fmt.Println(empire.PrettyPrintJson(starships))
	fmt.Println(empire.PrettyPrintJson(characters))
	fmt.Println(empire.PrettyPrintJson(planets))
}
