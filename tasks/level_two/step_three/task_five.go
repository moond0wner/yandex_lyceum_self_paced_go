// "Комплексные числа"

package main

import "fmt"

func PrintComplexNumber(z complex64) {
	fmt.Printf("%.2f %.2f", real(z), imag(z))
}
