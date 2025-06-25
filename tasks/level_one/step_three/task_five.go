// "Красивый текст"

package main

import (
	"fmt"
	"strings"
)

func main() {
	var numberOfWords, countFindTarget int
	var targetString string

	fmt.Scan(&numberOfWords)
	fmt.Scan(&targetString)

	lowerTargetString := strings.ToLower(targetString)

	for i := 0; i < numberOfWords; i++ {
		var input string
		fmt.Scan(&input)

		if strings.ToLower(input) == lowerTargetString {
			countFindTarget += 1
		}
	}
	fmt.Println(countFindTarget)

}
