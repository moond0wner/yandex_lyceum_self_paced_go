// "Go или не Go, вот в чём вопрос, v. 2.0"

package main

import (
	"fmt"
)

func main() {
	var input string

	for i := 0; i < 5; i++ {
		fmt.Scan(&input)
		if input == "Go" {
			fmt.Println("Go!")
		} else {
			fmt.Println("Я знаю только Go!")
		}
	}
}
