package main

import (
	"aluUpgradeCalc/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":8000"

	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/search", handlers.CarSearch)
	http.HandleFunc("/addCar", handlers.AddCar)
	http.HandleFunc("/removeCar", handlers.RemoveCar)
	http.HandleFunc("/calculate", handlers.PrintCars)

	fmt.Println("Listening on ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
