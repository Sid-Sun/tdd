package main

import (
	"errors"
	"math"
)

type square struct {
	side int
}

func NewSquare(side int) (square,error){
	if side < 1 {
		return square{},errors.New("Invalid size")
	}
	return square{side:side},nil
}

func (s square) area() int {
	return s.side * s.side
}

func (s square) perimeter() int {
	if math.Signbit(float64(s.side)) {
		return 0
	}
	return 4 * s.side
}
