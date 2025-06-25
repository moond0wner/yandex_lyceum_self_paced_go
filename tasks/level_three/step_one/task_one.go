// "Оптимизация памяти"

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func CountLengthAndBytes(first, second string) string {
	countRuneInFirst := utf8.RuneCountInString(first)
	countRuneInSecond := utf8.RuneCountInString(second)
	countByteInFirst := len(first)
	countByteInSecond := len(second)

	words := []string{first, second}

	return fmt.Sprintf(
		"Объединённая строка: %s. Количество байт: %d. Количество символов: %d.",
		strings.Join(words, ""),
		countByteInFirst+countByteInSecond,
		countRuneInFirst+countRuneInSecond,
	)
}
