// "Сортировка слиянием"

package main

import (
	"slices"
)

func sortNums(nums []int) []int {
	slices.SortFunc(nums, func(a, b int) int {
		return a - b
	})
	return nums
}

func SortAndMerge(left, right []int) []int {
	sortedLeftSlice := sortNums(left)
	sortedRightSlice := sortNums(right)
	mergeAndSortedSlice := append(sortedLeftSlice, sortedRightSlice...)
	return sortNums(mergeAndSortedSlice)
}
