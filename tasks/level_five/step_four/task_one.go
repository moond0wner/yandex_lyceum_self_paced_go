// "Самый-самый"

package main

import "slices"

func SortNums(nums []uint) {
	slices.SortFunc(nums, func(a, b uint) int {
		return int(a - b)
	})
}
