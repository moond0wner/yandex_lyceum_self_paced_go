// "Квадратные уравнения"

package main

import (
	"math"
)

func findDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4.0*a*c
}

func SquareRoots(a, b, c float64) (float64, float64) {
	discriminant := findDiscriminant(a, b, c)
	if discriminant > 0 {
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		return x2, x1
	} else if discriminant == 0 {
		x := -b / (2 * a)
		return x, x
	} else {
		return 0, 0
	}
}
