// "Банковский счёт"

package main

import "errors"

var (
	ErrBalanceIsNegative   = errors.New("Balance is negative")
	ErrValueIsNegative     = errors.New("Value is negative")
	ErrValueGreaterBalance = errors.New("The value is greater than the balance")
)

type Account struct {
	balance float64
	owner   string
}

func NewAccount(owner string) *Account {
	return &Account{balance: 0.0, owner: owner}
}

func (a *Account) SetBalance(value float64) error {
	if a.balance < 0 {
		return ErrBalanceIsNegative
	}
	if value < 0 {
		return ErrValueIsNegative
	}
	a.balance = value
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(value float64) error {
	if a.balance < 0 {
		return ErrBalanceIsNegative
	}
	if value < 0 {
		return ErrValueIsNegative
	}
	a.balance += value
	return nil
}

func (a *Account) Withdraw(value float64) error {
	if a.balance < 0 {
		return ErrBalanceIsNegative
	}
	if value < 0 {
		return ErrValueIsNegative
	}
	if a.balance < value {
		return ErrValueGreaterBalance
	}
	a.balance -= value
	return nil
}
