package main

import "fmt"

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

func (a Audi) Milaze() float64 {
	return a.distance / a.fuel
}
func totalMotorMilaz(m []MotorVhical) {
	tm := 0.0
	for _, v := range m {
		tm = tm + v.Milaze()
	}
	fmt.Printf("Your milaze is %f\n", tm)
}
func main() {
	b1 := Bmw{
		distance:     56.2,
		fuel:         52.3,
		averagespeed: "56.2",
	}
	a1 := Audi{
		distance: 64.3,
		fuel:     21.2,
	}
	persion := []MotorVhical{b1, a1}
	totalMotorMilaz(persion)
}
