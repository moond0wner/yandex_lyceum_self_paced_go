// "Слияние двух частей"

package main

func Mix(nums []int) []int {
	n := len(nums) / 2
	newSlice := make([]int, 0, len(nums))

	for i := 0; i < n; i++ {
		newSlice = append(newSlice, nums[i], nums[n+i])
	}

	return newSlice
}
