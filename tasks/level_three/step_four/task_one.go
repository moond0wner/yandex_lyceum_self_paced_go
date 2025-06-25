// "Виртуальный тайник"

package main

func FindMaxKey(n map[int]int) int {
	max := n[0]
	for num, _ := range n {
		if num > max {
			max = num
		}
	}
	return max
}
