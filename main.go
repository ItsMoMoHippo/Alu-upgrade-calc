package main

import (
	"aluUpgradeCalc/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":8000"
	router := http.DefaultServeMux

	router.HandleFunc("/", handlers.HandleHome)
	router.HandleFunc("/search", handlers.CarSearch)
	router.HandleFunc("/addCar", handlers.AddCar)
	router.HandleFunc("/removeCar", handlers.RemoveCar)
	router.HandleFunc("/calculate", handlers.PrintCars)

	fmt.Println("Listening on ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
