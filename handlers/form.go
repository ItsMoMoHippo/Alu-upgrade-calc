package handlers

import (
	"aluUpgradeCalc/data"
	"aluUpgradeCalc/database"
	"aluUpgradeCalc/views"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	log.Println("User added: ", car)
	_ = views.CarConfComponent(car, maxUpgrades).Render(r.Context(), w)
}

func SaveCar(w http.ResponseWriter, r *http.Request) {
	// TODO: parse data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad form", http.StatusBadRequest)
		return
	}

	car := data.CarConfSimple{
		Name: r.Form.Get("name"),
	}

	var err error
	if car.Speed, err = parseIntField(r, "speed"); err != nil {
		http.Error(w, "invalid speed stat", http.StatusBadRequest)
		return
	}
	if car.Accel, err = parseIntField(r, "accel"); err != nil {
		http.Error(w, "invalid speed stat", http.StatusBadRequest)
		return
	}
	if car.Handling, err = parseIntField(r, "handling"); err != nil {
		http.Error(w, "invalid speed stat", http.StatusBadRequest)
		return
	}
	if car.Nitro, err = parseIntField(r, "nitro"); err != nil {
		http.Error(w, "invalid speed stat", http.StatusBadRequest)
		return
	}

	_ = views.CarSavedComp(car).Render(r.Context(), w)
}

func RemoveCar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	var cars []data.CarConfSimple
	if err := json.NewDecoder(r.Body).Decode(&cars); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// caluclate
}

func parseIntField(r *http.Request, field string) (int, error) {
	val := r.Form.Get(field)
	if val == "" {
		return 0, nil
	}

	return strconv.Atoi(val)
}
