/*
Package data: Contains all data and structs needed and database connectivity
*/
package data

import (
	_ "embed"
	"errors"
)

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

func MaxUpgradeStage(star int) (int, error) {
	switch star {
	case 3:
		return 10, nil
	case 4:
		return 11, nil
	case 5:
		return 12, nil
	case 6:
		return 13, nil
	}
	return 0, errors.New("invalid number")
}
