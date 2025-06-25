// "Сортировка подсчётом"

package main

func CountingSort(contacts []string) map[string]int {
	resultMap := map[string]int{}
	for _, element := range contacts {
		resultMap[element]++
	}
	return resultMap
}
