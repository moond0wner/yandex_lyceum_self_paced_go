// "Кукушка"

package main

import (
	"fmt"
	"time"
)

func main() {
	var timeInput string
	var layout string = "2006-01-02/15:04:05"
	fmt.Scanln(&timeInput)
	parsedTime, err := time.Parse(layout, timeInput)

	if err != nil {
		fmt.Println("Ошибка при разборе: ", err)
		return

	}

	hours := parsedTime.Hour()
	minutes := parsedTime.Minute()

	fmt.Println(
		fmt.Sprintf(
			"Текущее время %d часов, %d минут. Ты точно не забыл про важные дела на сегодня?",
			hours,
			minutes,
		),
	)
}
