// "Пароли возвращение"

package main

import (
	"unicode"
)

// uncomment!!!
// func hasMinimumLength(password string, minimalLengthPassword int) bool {
// 	return len(password) >= minimalLengthPassword
// }

// func hasUpper(password string) bool {
// 	for _, letter := range password {
// 		if unicode.IsUpper(letter) {
// 			return true
// 		}
// 	}
// 	return false
// }

func hasLowerCase(password string) bool {
	for _, letter := range password {
		if unicode.IsLower(letter) {
			return true
		}
	}
	return false
}

func ratePassword(password string) string {
	var rating int
	if hasLowerCase(password) {
		rating++
	}
	if hasMinimumLength(password, 8) {
		rating++
	}
	if hasUpper(password) {
		rating++
	}

	switch rating {
	case 1:
		return "Слабый пароль"
	case 2:
		return "Средний пароль"
	case 3:
		return "Сложный пароль"
	default:
		return "Пароль недостаточно безопасен, придумайте новый."
	}
}
