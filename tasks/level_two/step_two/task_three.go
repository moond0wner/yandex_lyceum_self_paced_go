// "Пароли"

package main

import (
	"unicode"
)

func hasMinimumLength(password string, minimalLengthPassword int) bool {
	return len(password) >= minimalLengthPassword
}

func hasUpper(password string) bool {
	for _, letter := range password {
		if unicode.IsUpper(letter) {
			return true
		}
	}
	return false
}

func checkPassword(password string) bool {
	return hasMinimumLength(password, 8) && hasUpper(password)
}
