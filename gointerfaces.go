/*
Summary:
Create a program that directs vehicles at a mechanic shop
to the correct vehicle lift, based on vehicle size

Requirements:
The shop has lifts for multiple vehicle sizes/types:
- Motorcycles: small lifts
- Cars: standard lifts
- Trucks: large lifts

Write a single function to handle all of the vehicles
that the shop works on.

Vehicles have a model name in addition to the vehicle type:
- Example: "Truck" is the vehicle type, "Road Devourer" is a model name

Direct at least 1 of each vehicle type to the correct
lift, and print out the vehicle information.

Notes:
Use any names for vehicle models
*/

package main

import (
	"fmt"
)

type LiftPicker interface {
	PickLift() Lift
}
type Lift int

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) PickLift() Lift {
	return SmallLift
}

func (m Car) PickLift() Lift {
	return StandardLift
}

func (m Truck) PickLift() Lift {
	return LargeLift
}

func DirectToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		fmt.Printf("%v directed to small lift\n", p)
	case StandardLift:
		fmt.Printf("%v directed to standard lift\n", p)
	case LargeLift:
		fmt.Printf("%v directed to large lift\n", p)
	}
}

func main() {
	m := Motorcycle("Harley Davidson")
	c := Car("Honda Civic")
	t := Truck("Toyota Tundra")
	DirectToLift(m)
	DirectToLift(c)
	DirectToLift(t)
}
