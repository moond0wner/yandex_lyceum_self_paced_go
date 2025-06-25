// "Буква ё"

package main

import (
	"fmt"
	"strings"
)

func checkLetters(text string) string {
	countTargetLetter := strings.Count(text, "е")
	if countTargetLetter > 0 {
		return fmt.Sprintf("Количество возможных ошибок: %d, перепроверьте текст")
	}
	return "Текст готов к публикации!"
}
