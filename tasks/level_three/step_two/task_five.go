// "Сделал дело - гуляй смело"

package main

import "fmt"

func PrettyArrayOutput(array [9]string) {
	i := 1
	for i <= 7 {
		fmt.Printf("%d я уже сделал: %s\n", i, array[i-1])
		i++
	}
	fmt.Printf("%d не успел сделать: %s\n", i, array[i-1])
	i++
	fmt.Printf("%d не успел сделать: %s\n", i, array[i-1])
}
