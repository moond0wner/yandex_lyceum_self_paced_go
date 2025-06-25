// "Инверсия ключей и значений"

package main

func SwapKeysAndValues(m map[string]string) map[string]string {
	resultMap := map[string]string{}
	for key, value := range m {
		resultMap[value] = key
	}
	return resultMap
}
