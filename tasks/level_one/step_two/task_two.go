// "Надежный пароль"

package main

import (
	"fmt"
)

func isReliablePassword(pass string) bool {
	return len(pass) >= 8
}

func TaskTwo() {
	var passOne, passTwo string
	fmt.Scan(&passOne)
	fmt.Scan(&passTwo)

	if isReliablePassword(passOne) && isReliablePassword(passTwo) {
		fmt.Println("Оба пароля надежные")
	} else if isReliablePassword(passOne) {
		fmt.Println("Только первый пароль надежный")
	} else if isReliablePassword(passTwo) {
		fmt.Println("Только второй пароль надежный")
	} else {
		fmt.Println("Оба пароля ненадёжные")
	}
}
