// "Чеки"

package main

import (
	"fmt"
)

func main() {
	var numberOfProducts, discount int
	var sumOfPricesWithDiscount float64
	fmt.Scan(&numberOfProducts)
	fmt.Scan(&discount)
	for i := 0; i < numberOfProducts; i++ {
		var priceProduct float64
		fmt.Scan(&priceProduct)

		discountAmount := priceProduct * (float64(discount) / 100.0)
		priceWithDiscount := priceProduct - discountAmount

		sumOfPricesWithDiscount += priceWithDiscount
	}

	fmt.Println(sumOfPricesWithDiscount)
}
