// "Табло аэропорта"

package main

import (
	"fmt"
)

func PrintFlightRow(numberFlight string, cityStart string, cityEnd string, duration float64, numberReception int, numberGate int, registrationDone bool) {
	if registrationDone {
		fmt.Printf(
			"| %s | %s—%s | регистрация закончилась, проходите к гейту: %d | длительность полёта %.1f часа |\n",
			numberFlight,
			cityStart,
			cityEnd,
			numberGate,
			duration,
		)
	} else {
		fmt.Printf(
			"| %s | %s—%s | %d регистрация продолжается |\n",
			numberFlight,
			cityStart,
			cityEnd,
			numberReception,
		)
	}
}
