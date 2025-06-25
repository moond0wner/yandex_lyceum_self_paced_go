// "Бесконечность не предел"

package main

import (
	"fmt"
	"math"
)

func main() {
	var numFirst, numSecond float64

	fmt.Scanln(&numFirst)
	fmt.Scanln(&numSecond)

	fmt.Println(math.Floor(math.Max(numFirst, numSecond)))
}
