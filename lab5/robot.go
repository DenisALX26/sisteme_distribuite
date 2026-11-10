package main

import "math"

type Point struct {
	X, Y float64
}

type Robot struct {
	Position Point
	Activ    bool
}

func rotunjire(val float64) float64 {
	return math.Round(val*10000) / 10000
}
