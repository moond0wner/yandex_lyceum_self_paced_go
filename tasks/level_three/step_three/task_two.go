// "Слияние"

package main

func Join(nums1, nums2 []int) []int {
	newSlice := make([]int, len(nums1)+len(nums2), len(nums1)+len(nums2))
	for i := 0; i < len(nums1); i++ {
		newSlice[i] = nums1[i]
	}
	for i := 0; i < len(nums2); i++ {
		newSlice[len(nums1)+i] = nums2[i]
	}
	return newSlice
}
