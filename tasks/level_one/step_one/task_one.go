// "Go или Go, вот в чём вопрос"

package main

import "fmt"

func main() {
	var stroka string
	fmt.Scanln(&stroka)

	if stroka == "Go" {
		fmt.Println("Go!")
		return
	}
	fmt.Println("Я знаю только Go!")
}
