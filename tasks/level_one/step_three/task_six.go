// "Секретный язык"

package main

import (
	"fmt"
)

func main() {
	var inputWord, resultWord string

	fmt.Scan(&inputWord)

	for _, letter := range inputWord {
		if string(letter) == "а" {
			continue
		} else if string(letter) == "о" {
			continue
		}
		resultWord += string(letter)
	}
	fmt.Println(resultWord)
}
