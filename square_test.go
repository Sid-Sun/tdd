package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSquareArea(t *testing.T) {
	four, _ := NewSquare(4)
	result := four.area()
	assert.Equal(t, 16, result)
	nine, _ := NewSquare(9)
	resultTwo := nine.area()
	assert.Equal(t, 81, resultTwo)
}

func TestSquarePerimeter(t *testing.T) {
	five, _ := NewSquare(5)
	result := five.perimeter()
	assert.Equal(t, 20, result)
	four, _ := NewSquare(4)
	resultTwo := four.perimeter()
	assert.Equal(t, 16, resultTwo)
	_, err := NewSquare(-4)
	require.Error(t, err)
	//resultThree := NegativeFour.perimeter()
	//assert.Equal(t, 0, resultThree)
}
