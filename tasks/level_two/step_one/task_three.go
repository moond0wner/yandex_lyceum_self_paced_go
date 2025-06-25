// "Фастфуд"

package main

import (
	"fmt"
)

func BuyFries(size string) {
	switch size {
	case "S":
		fmt.Println("Картошка фри будет стоить 49 рублей")
	case "M":
		fmt.Println("Картошка фри будет стоить 89 рублей")
	case "L":
		fmt.Println("Картошка фри будет стоить 99 рублей")
	default:
		fmt.Println("Некорректный размер")
	}
}

func BuyCola(size string) {
	switch size {
	case "S":
		fmt.Println("Кола будет стоить 99 рублей")
	case "M":
		fmt.Println("Кола будет стоить 119 рублей")
	case "L":
		fmt.Println("Кола будет стоить 139 рублей")
	default:
		fmt.Println("Некорректный размер")
	}
}

func printPrice(price int, nameProduct string) {
	fmt.Printf("%s будет стоить %d рублей\n", nameProduct, price)
}
