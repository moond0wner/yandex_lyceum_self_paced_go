// "Путешествие по времени"

package main

import (
	"fmt"
	"time"
)

func main() {
	var layout string = "2006-01-02"
	var future, present string

	fmt.Scanln(&future)
	fmt.Scanln(&present)

	futureTime, futureTimeError := time.Parse(layout, future)
	presentTime, presentTimeError := time.Parse(layout, present)

	if futureTimeError != nil {
		fmt.Println("Ошибка:", futureTimeError)
		return
	}

	if presentTimeError != nil {
		fmt.Println("Ошибка:", presentTimeError)
		return
	}

	diff := futureTime.Year() - presentTime.Year()

	fmt.Println(
		fmt.Sprintf(
			"%d year ago",
			diff,
		),
	)
}
