// "Тест DeleteVowels"

package main

import (
	"testing"
)

type Test struct {
	in  string
	out string
}

var tests = []Test{
	{"AEIOU", ""},                     // Все гласные в верхнем регистре
	{"aeiou", ""},                     // Все гласные в нижнем регистре
	{"12345", "12345"},                // Числа
	{"!@#$%", "!@#$%"},                // Специальные символы
	{"mixed CASE Test", "mxd CS Tst"}, // Смешанный регистр
	{"", ""},                          // Пустая строка
	{"A B C D E", " B C D "},          // Пробелы и гласные
	{"Gonna go!", "Gnn g!"},           // Смешанные символы
}

func TestDeleteVowels(t *testing.T) {
	for i, test := range tests {
		got := DeleteVowels(test.in)
		if got != test.out {
			t.Errorf("#%d: DeleteVowels(%s)=%s; want %s", i, test.in, got, test.out)
		}
	}
}
