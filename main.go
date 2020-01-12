package main

import (
	"log"
	"net/http"
	"swapi-go-db/empire"
)

func main() {
	empire.ConnectDB()
	server := empire.NewStarWarsServer()

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalln(err)
	}
}
