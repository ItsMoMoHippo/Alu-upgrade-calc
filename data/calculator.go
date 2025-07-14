package data

func CalculateCars(cars []CarConfig) int {
	total := 0
	for _, car := range cars {
		total += CalculateCar(car)
	}
	return total
}

func CalculateCar(CarConfig) int {
	return 1
}

func GetStats() (CarConfig, error) {
	// unmarshall object from json array from user
	name := "foo"
	ss := 1
	as := 1
	hs := 1
	ns := 1
	cip := 1
	rip := 1
	eip := 1

	return CarConfig{
		Name:          name,
		SpeedStage:    ss,
		AccelStage:    as,
		HandlingStage: hs,
		NitroStage:    ns,
		CommImpParts:  cip,
		RareImpParts:  rip,
		EpicImpParts:  eip,
	}, nil
}
