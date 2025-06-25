// "Тест Sum"

package main

import "testing"

func TestSum(t *testing.T) {
	got := Sum(2, 2)
	excepted := 4
	if got != excepted {
		t.Fatalf(`Sum(a, b) = %d, want %d`, got, excepted)
	}
}
