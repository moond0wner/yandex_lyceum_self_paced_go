// "NFT-артефакты"

package main

func SumOfValuesInMap(m map[int]int) int {
	var sumOfValues int
	for _, count := range m {
		sumOfValues += count
	}
	return sumOfValues
}
