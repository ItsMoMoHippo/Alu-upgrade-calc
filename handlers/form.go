package handlers

import (
	"aluUpgradeCalc/data"
	"aluUpgradeCalc/views"
	"fmt"
	"net/http"
	"strings"

	"github.com/a-h/templ"
)

func CarSearch(w http.ResponseWriter, r *http.Request) {
	frag := strings.ToLower(r.URL.Query().Get("carInput"))
	var matches []string
	for _, car := range data.CarsList {
		if strings.Contains(strings.ToLower(car), frag) {
			matches = append(matches, car)
		}
	}
	templ.Handler(views.CarOptions(matches)).ServeHTTP(w, r)
}

func PrintCars(w http.ResponseWriter, r *http.Request) {
	cars := r.Form["cars"]
	fmt.Printf("cars recieved: %v\n", cars)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Println(w, "sent %d cars", len(cars))
}

func AddCar(w http.ResponseWriter, r *http.Request) {
	car := strings.TrimSpace(r.FormValue("carInput"))
	if car == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_ = views.CarRow(car).Render(r.Context(), w)
}

func RemoveCar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
