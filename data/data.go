/*
Package data: Contains all data and structs needed
*/
package data

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed cars.json
var carsJSON []byte

var CarsList []string

func init() {
	err := json.Unmarshal(carsJSON, &CarsList)
	if err != nil {
		log.Fatal("Failed to unmarshall cars.json : %w", err)
	}
}

type CarConfig struct {
	Name          string `json:"name"`
	SpeedStage    int    `json:"speed"`
	AccelStage    int    `json:"accel"`
	HandlingStage int    `json:"handling"`
	NitroStage    int    `json:"nitro"`
	CommImpParts  int    `json:"cip"`
	RareImpParts  int    `json:"rip"`
	EpicImpParts  int    `json:"eip"`
}
