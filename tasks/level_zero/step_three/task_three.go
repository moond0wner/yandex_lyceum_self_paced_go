// "Приглашения"

package main

import "fmt"

func main() {
	var nameFriend string
	var houseNumber int
	var secretNumber int
	var durationCompetion float64

	fmt.Scanln(&nameFriend)
	fmt.Scanln(&houseNumber)
	fmt.Scanln(&secretNumber)
	fmt.Scanln(&durationCompetion)

	invitationLetter := fmt.Sprintf(
		"Привет, %s! Приглашаю тебя на соревнование по программированию, которое пройдёт, как всегда, в квартире %d. Оно будет длиться примерно %.1f часа. Не забудь секретный пароль для входа: %d.",
		nameFriend,
		houseNumber,
		durationCompetion,
		secretNumber,
	)

	fmt.Println(invitationLetter)

}
