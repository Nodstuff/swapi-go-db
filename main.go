package main

import (
	"encoding/json"
	"net/http"
	"swapi-go-db/empire"
)

func main() {
	http.HandleFunc("/", filmHandler)
	http.ListenAndServe(":8080", nil)
}

func filmHandler(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(empire.GetFilmsSql(empire.ConnectSql()))
	empire.CheckErr(err)
	w.Write(j)
}
