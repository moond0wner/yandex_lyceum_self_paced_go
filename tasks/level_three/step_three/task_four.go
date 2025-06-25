// "Неизбирательный шоппинг"

package main

import (
	"errors"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	result := []int{}
	if n < 0 {
		return nil, errors.New("n отрицательное")
	}
	if nums == nil {
		return nil, errors.New("слайс пустой")
	}
	for _, num := range nums {
		if num < limit {
			result = append(result, num)
			if len(result) == n {
				break
			}
		}
	}
	return result, nil
}
