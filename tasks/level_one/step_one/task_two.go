// "Субъективно больше"

package main

import "fmt"

func main() {
	var numberOne, numberTwo int
	fmt.Scan(&numberOne)
	fmt.Scan(&numberTwo)

	if numberOne > numberTwo {
		fmt.Println("Первое число больше второго")
	} else if numberOne < numberTwo {
		fmt.Println("Второе число больше первого")
	} else {
		fmt.Println("Числа равны")
	}
}
