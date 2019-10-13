package main

import "errors"

type Money struct {
	Amount int
}

func NewMoney(value int) (Money, error) {
	if value<0 {
		return Money{}, errors.New("MONEY CANNOT BE LESS THAN 0")
	}
	return Money{Amount: value}, nil
}

func (MyMoney Money) Add(YourMoney Money) Money {
	return Money{MyMoney.Amount + YourMoney.Amount}
}
