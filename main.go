package main

import (
	"database/sql"
	"log"
	"net/http"
	"swapi-go-db/empire"
	"swapi-go-db/empire/database"
)

func main() {
	var err error
	database.DBCon, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/sw_data")
	empire.CheckErr(err)

	server := empire.NewStarWarsServer()

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalln(err)
	}
}

