// "Глобальная чистка"

package main

func DeleteLongKeys(m map[string]int) map[string]int {
	resultMap := map[string]int{}
	for contact, number := range m {
		if len(contact) >= 6 {
			resultMap[contact] = number
		}
	}
	return resultMap
}
