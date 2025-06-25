package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var layout string = "02.01.2006"
	var dateStartJob, name, surname, patronymic string
	var sumOne, sumTwo, sumThree float64

	fmt.Scanln(&dateStartJob)
	fmt.Scanln(&name)
	fmt.Scanln(&surname)
	fmt.Scanln(&patronymic)
	fmt.Scanln(&sumOne)
	fmt.Scanln(&sumTwo)
	fmt.Scanln(&sumThree)

	parsedTime, err := time.Parse(layout, dateStartJob)

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	futureDate := parsedTime.AddDate(0, 0, 15).Format(layout)
	totalSum := sumOne + sumTwo + sumThree
	rubles := int(totalSum)
	kopecks := int(math.Round(math.Mod(totalSum*100, 100)))

	fmt.Println(
		fmt.Sprintf(
			"Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы. \nДата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день. \nОбщая сумма выплат составит %d руб. %d коп. \n\nС уважением, \nГл. бух. Иванов А.Е.",
			surname,
			name,
			patronymic,
			futureDate,
			rubles,
			kopecks,
		),
	)
}
