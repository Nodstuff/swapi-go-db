package main

import (
	"log"
	"net/http"
	"swapi-go-db/empire"
)

func main() {
	empire.ConnectDB()
	log.Fatalln(http.ListenAndServe(":8080", empire.NewStarWarsServer()))
}
