package main

import (
	"aluUpgradeCalc/handlers"
	"log"
	"net/http"
)

func main() {
	port := ":8000"
	router := http.DefaultServeMux

	router.HandleFunc("/", handlers.HandleHome)
	router.HandleFunc("/search", handlers.CarSearch)
	router.HandleFunc("/addCar", handlers.AddCar)
	router.HandleFunc("/saveCar", handlers.SaveCar)
	router.HandleFunc("/removeCar", handlers.RemoveCar)
	router.HandleFunc("/calculate", handlers.PrintCars)

	log.Print("Listening on ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
