// "Покупки"

package main

func SumOfArray(array [6]int) int {
	var sumOfArray int
	for _, num := range array {
		sumOfArray += num
	}
	return sumOfArray
}
