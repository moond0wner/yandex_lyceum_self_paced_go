// "Интересные числа"

package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64
	fmt.Scanln(&number)
	result := math.Sqrt(number)
	if !math.IsNaN(result) {
		fmt.Printf("%.3f", result)
		return
	}
	fmt.Println(-1)

}
