package main

type MotorVhical interface {
	Milaze() float64
}

type Bmw struct {
	distance     float64
	fuel         float64
	averagespeed string
}

type Audi struct {
	distance float64
	fuel     float64
}

func (b Bmw) Milaze() float64 {
	return b.distance / b.fuel
}
