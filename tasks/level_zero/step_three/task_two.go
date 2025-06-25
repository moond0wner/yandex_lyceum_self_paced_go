// "Вот что я люблю"

package main

import "fmt"

func main() {
	var enteredValue string
	fmt.Scanln(&enteredValue)
	fmt.Println(
		fmt.Sprintf(
			"А ещё я люблю %s!",
			enteredValue,
		),
	)
}
