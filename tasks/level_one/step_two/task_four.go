// "Модный приговор"

package main

import (
	"fmt"
)

func TaskFour() {
	var sign string
	var number int
	fmt.Scan(&sign)
	if sign == "0" {
		number = 0
	} else {
		fmt.Scan(&number)
	}

	switch {
	case sign == "+" && number > 20:
		fmt.Println("Стоит надеть майку и шорты")
	case sign == "+" && number >= 10 && number <= 20:
		fmt.Println("Стоит надеть штаны и кофту")
	case (sign == "+" && number >= -5 && number <= 9) || (sign == "-" && number >= -5 && number <= 9) || (sign == "0" && number == 0):
		fmt.Println("Стоит надеть куртку")
	default:
		fmt.Println("Стоит надеть зимнюю куртку")
	}

}
