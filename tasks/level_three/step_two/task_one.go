// "Пять шагов"

package main

func FiveSteps(array [5]int) [5]int {
	var newArray [5]int = [5]int{array[4], array[3], array[2], array[1], array[0]}
	return newArray
}
