// "Камень, ножницы и бумага"

package main

import (
	"fmt"
)

func TaskThree() {
	var resultOne, resultTwo string
	fmt.Scan(&resultOne)
	fmt.Scan(&resultTwo)

	switch {
	case resultOne == resultTwo:
		fmt.Println("Ничья")
	case resultOne == "камень" && resultTwo == "ножницы":
		fmt.Println("Первый игрок победил")
	case resultOne == "ножницы" && resultTwo == "бумага":
		fmt.Println("Первый игрок победил")
	case resultOne == "бумага" && resultTwo == "камень":
		fmt.Println("Первый игрок победил")
	default:
		fmt.Println("Второй игрок победил")
	}

}
