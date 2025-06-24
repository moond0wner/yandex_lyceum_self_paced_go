package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

// Так как TimeNow реализована в LMS, то для работы кода напишу заглушку
func TimeNow() time.Time {
	return time.Now()
}

var ErrGoSleep = errors.New("исправь свой ответ, а лучше ложись поспать")

func currentDayOfTheWeek() string {
	var day string
	switch TimeNow().Weekday() {
	case time.Monday:
		day = "Понедельник"
	case time.Tuesday:
		day = "Вторник"
	case time.Wednesday:
		day = "Среда"
	case time.Thursday:
		day = "Четверг"
	case time.Friday:
		day = "Пятница"
	case time.Saturday:
		day = "Суббота"
	case time.Sunday:
		day = "Воскресенье"
	}
	return day
}

func dayOrNight() string {
	currentHour := TimeNow().Hour()
	if currentHour >= 10 && currentHour <= 22 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
	today := TimeNow()
	nextFriday := (time.Friday - today.Weekday() + 7) % 7
	if nextFriday == 0 {
		return 0
	}
	return int(nextFriday)
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	answer = strings.ToLower(answer)
	currentDay := strings.ToLower(currentDayOfTheWeek())
	return answer == currentDay
}

func CheckNowDayOrNight(answer string) (bool, error) {
	currentTime := strings.ToLower(dayOrNight())
	length := utf8.RuneCountInString(answer)

	answer = strings.ToLower(answer)
	if length > 4 || length < 4 {
		return false, ErrGoSleep
	}
	if answer == currentTime {
		return true, nil
	}
	return false, nil
}
