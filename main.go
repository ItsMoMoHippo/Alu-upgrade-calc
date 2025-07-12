package main

import (
	"aluUpgradeCalc/views"
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	port := ":8000"

	http.Handle("/", templ.Handler(views.HomePage("Asphalt Legends Unite - Upgrade Calculator")))

	fmt.Println("Listening on ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
