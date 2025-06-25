// "Абсолютное копирование"

package main

func SliceCopy(nums []int) []int {
	newSlice := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		newSlice[i] = nums[i]
	}
	return newSlice
}
