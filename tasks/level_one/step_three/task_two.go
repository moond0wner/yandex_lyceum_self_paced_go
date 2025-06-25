// "Экзамен"

package main

import (
	"fmt"
)

func main() {
	var numberOfStudents int
	fmt.Scan(&numberOfStudents)

	for i := 0; i < numberOfStudents; i++ {
		var score float64
		fmt.Scan(&score)

		switch {
		case score >= 0 && score <= 49:
			fmt.Println(2)
		case score >= 50 && score <= 74:
			fmt.Println(3)
		case score >= 75 && score <= 89:
			fmt.Println(4)
		case score >= 90 && score <= 100:
			fmt.Println(5)
		default:
			fmt.Println("Неверный балл")
		}
	}
}
