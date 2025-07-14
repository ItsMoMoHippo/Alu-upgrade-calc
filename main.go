package main

import (
	"aluUpgradeCalc/views"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/a-h/templ"
)

//go:embed cars.json
var carsJSON []byte

var carsList []string

func init() {
	err := json.Unmarshal(carsJSON, &carsList)
	if err != nil {
		log.Fatal("Failed to unmarshall cars.json : %w", err)
	}
}

func main() {
	port := ":8000"

	http.Handle("/", templ.Handler(views.HomePage("Asphalt Legends Unite - Upgrade Calculator", []string{})))
	http.HandleFunc("/search", CarSearch)

	fmt.Println("Listening on ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func CarSearch(w http.ResponseWriter, r *http.Request) {
	frag := strings.ToLower(r.URL.Query().Get("carInput"))
	var matches []string
	for _, car := range carsList {
		if strings.Contains(strings.ToLower(car), frag) {
			matches = append(matches, car)
		}
	}
	templ.Handler(views.CarOptions(matches)).ServeHTTP(w, r)
}
