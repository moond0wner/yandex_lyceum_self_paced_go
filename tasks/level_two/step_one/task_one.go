// "Go или не Go, вот в чём процедурный вопрос"

package main

import (
	"fmt"
)

func GoOrNot(word string) {
	if word == "Go" {
		fmt.Println("Go!")
	} else {
		fmt.Println("Я знаю только Go!")
	}
}
