// "Максимальное равество"

package main

import (
	"fmt"
)

func TaskOne() {
	var numOne, numTwo, numThree float64
	fmt.Scanln(&numOne, &numTwo, &numThree)

	if numOne == numTwo && numTwo == numThree {
		fmt.Println("Максимальное равенство")
	} else {
		fmt.Println("Не равны")
	}
}
