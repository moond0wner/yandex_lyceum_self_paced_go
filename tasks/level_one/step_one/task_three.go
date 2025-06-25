// "Вкусный кофе"

package main

import "fmt"

func main() {
	var rubles, kopecks, priceCoffee int
	fmt.Scan(&rubles)
	fmt.Scan(&kopecks)
	fmt.Scan(&priceCoffee)

	rubles += kopecks / 100

	if rubles >= priceCoffee {
		fmt.Println("Сегодня будет вкусный кофе!")
		return
	}
	fmt.Println("Стоит подкопить")
}
