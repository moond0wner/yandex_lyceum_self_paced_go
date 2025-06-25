// "Текстовые вычисления"

package main

import (
	"strings"
)

func NumbersToLetters(input string) string {
	for _, letter_code := range input {
		letter := string(letter_code)
		switch letter {
		case "0":
			input = strings.ReplaceAll(input, "0", "ноль")
		case "1":
			input = strings.ReplaceAll(input, "1", "один")
		case "2":
			input = strings.ReplaceAll(input, "2", "два")
		case "3":
			input = strings.ReplaceAll(input, "3", "три")
		case "4":
			input = strings.ReplaceAll(input, "4", "четыре")
		case "5":
			input = strings.ReplaceAll(input, "5", "пять")
		case "6":
			input = strings.ReplaceAll(input, "6", "шесть")
		case "7":
			input = strings.ReplaceAll(input, "7", "семь")
		case "8":
			input = strings.ReplaceAll(input, "8", "восемь")
		case "9":
			input = strings.ReplaceAll(input, "9", "девять")
		case "+":
			input = strings.ReplaceAll(input, "+", "плюс")
		case "-":
			input = strings.ReplaceAll(input, "-", "минус")
		case "*":
			input = strings.ReplaceAll(input, "*", "умножить на")
		case "/":
			input = strings.ReplaceAll(input, "/", "разделить на")
		}
	}
	return input
}
