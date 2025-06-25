// "Абвгд"

package main

import (
	"slices"
)

func SortNames(names []string) {
	slices.SortFunc(names, func(a, b string) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
}
