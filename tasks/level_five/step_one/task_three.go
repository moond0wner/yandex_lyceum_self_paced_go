// "Тест Multiply"

package main

import "testing"

type Test struct {
	firstOperand  int
	secondOperand int
	out           int
}

var tests = []Test{
	{1, 2, 2},
	{2, 2, 4},
	{6, 6, 36},
	{5, 5, 25},
	{10, 10, 100},
}

func TestMultiply(t *testing.T) {
	for _, test := range tests {
		got := Multiply(test.firstOperand, test.secondOperand)
		if got != test.out {
			t.Fatalf(`Multiply(%d, %d) = %d, want %d`, test.firstOperand, test.secondOperand, got, test.firstOperand*test.secondOperand)
		}
	}
}
