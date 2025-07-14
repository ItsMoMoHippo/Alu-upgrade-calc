package main

type Class string

const (
	ClassS Class = "S"
	ClassA Class = "A"
	ClassB Class = "B"
	ClassC Class = "C"
	ClassD Class = "D"
)

type Car struct {
	Name     string
	Class    Class
	CurStars int
	MaxStars int
	Image    string
}
