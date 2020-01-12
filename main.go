package main

import (
	"log"
	"net/http"
	"swapi-go-db/empire"
)

func main() {
	log.Fatalln(http.ListenAndServe(":8080", empire.NewStarWarsServer()))
}
