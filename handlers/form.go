package handlers

import (
	"aluUpgradeCalc/data"
	"aluUpgradeCalc/database"
	"aluUpgradeCalc/views"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/a-h/templ"
)

func CarSearch(w http.ResponseWriter, r *http.Request) {
	frag := strings.ToLower(r.URL.Query().Get("carInput"))

	query := "SELECT name FROM cars WHERE LOWER(name) LIKE ?"
	rows, err := database.DB.Query(query, "%"+frag+"%")
	if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows:%v", err)
		}
	}()

	var matches []string
	for rows.Next() {
		var car string
		if err := rows.Scan(&car); err != nil {
			continue
		}
		matches = append(matches, car)
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
	var star int
	query := "SELECT star FROM cars WHERE LOWER(name) = LOWER(?)"
	row := database.DB.QueryRow(query, car)
	err := row.Scan(&star)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Database error: car not found", http.StatusInternalServerError)
			return
		}
		http.Error(w, "Database Error, star not found", http.StatusInternalServerError)
		return
	}
	maxUpgrades, _ := data.MaxUpgradeStage(star)
	_ = views.CarConfComponent(car, maxUpgrades).Render(r.Context(), w)
}

func RemoveCar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
