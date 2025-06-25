// "Вы поедете на бал?"

package main

import (
	"fmt"
)

func main() {
	var input string
	for {
		fmt.Scan(&input)
		switch input {
		case "да":
			fmt.Println("Поражение")
			return
		case "нет":
			fmt.Println("Поражение")
			return
		case "чёрный":
			fmt.Println("Поражение")
			return
		case "белый":
			fmt.Println("Поражение")
			return
		default:
			fmt.Println("Игра продолжается")
		}
	}
}
