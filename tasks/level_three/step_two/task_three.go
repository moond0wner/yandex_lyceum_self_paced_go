// "Случайность по максимуму"

package main

func FindMinMaxInArray(array [10]int) (int, int) {
	min := 10000
	max := -10000

	for _, num := range array {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return max, min
}
