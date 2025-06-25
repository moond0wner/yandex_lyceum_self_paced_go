// "Минимальный минимум"

package main

import (
	"fmt"
	"math"
)

func main() {
	var arseniyHeight float64
	var goshaHeight float64
	var iraHeight float64

	fmt.Scanln(&arseniyHeight)
	fmt.Scanln(&goshaHeight)
	fmt.Scanln(&iraHeight)

	minimum := math.Min(math.Min(goshaHeight, arseniyHeight), iraHeight)
	fmt.Println(minimum)
}
