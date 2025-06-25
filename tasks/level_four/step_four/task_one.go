// "Система управления транспортом"

package main

import (
	"reflect"
)

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
}

type Car struct {
	Type     string
	Speed    float64
	FuelType string
}

type Motorcycle struct {
	Type     string
	Speed    float64
	FuelType string
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / m.Speed
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	travelTimes := map[string]float64{}
	for i := 0; i < len(vehicles); i++ {
		travelTime := vehicles[i].CalculateTravelTime(distance)
		travelTimes[reflect.TypeOf(vehicles[i]).String()] = travelTime
	}
	return travelTimes
}
