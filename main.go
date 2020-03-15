package main

import (
	"fmt"
	"log"
	"net/http"
	"swapi-go-db/empire"
)

func main() {
	fmt.Println("Starting server")
	log.Fatalln(http.ListenAndServe(":8100", empire.NewStarWarsServer()))
}
