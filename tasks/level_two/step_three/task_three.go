// "Бесконечные траты"

package main

import (
	"errors"
)

var (
	Balance      int
	AmountError  = errors.New("amount is incorrect")
	BalanceError = errors.New("balance is incorrect")
)

func topUpBalance(amount float64) error {
	if amount <= 0 {
		return AmountError
	}
	Balance += int(amount)
	return nil
}

func chargeFromBalance(amount float64) error {
	if amount <= 0 {
		return AmountError
	}
	if Balance < int(amount) {
		return BalanceError
	}
	Balance -= int(amount)
	return nil
}

func TopUpAndGetBalance(amount float64) (float64, error) {
	err := topUpBalance(amount)
	if err != nil {
		return 0, err
	}
	if Balance < 0 {
		return 0, BalanceError
	}
	return float64(Balance), nil
}

func ChargeFromAndGetBalance(amount float64) (float64, error) {
	err := chargeFromBalance(amount)
	if err != nil {
		return 0, err
	}
	if Balance < 0 {
		return 0, BalanceError
	}
	return float64(Balance), nil
}
