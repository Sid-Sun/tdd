package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewProduct(t *testing.T) {
	apple, err := NewProduct("Apple", 12)
	require.NoError(t, err)
	assert.Equal(t, 12, apple.price)
	assert.Equal(t, "Apple", apple.name)
	_, err = NewProduct("", 12)
	require.Error(t, err)
	_, err = NewProduct("", -78)
	require.Error(t, err)
	_, err = NewProduct("Apple", 0)
	require.Error(t, err)
	_, err = NewProduct("Apple", -54)
	require.Error(t, err)
}

func TestNewCart(t *testing.T) {
	cart := NewCart()
	assert.Equal(t, Cart{}, cart)
}

func TestCart_Add(t *testing.T) {
	apple, _ := NewProduct("Apple", 12)
	banana, _ := NewProduct("Banana", 18)
	cart := NewCart()
	cart = cart.Add(apple)
	assert.Equal(t, apple, cart.products[0])
	assert.Equal(t, 1, len(cart.products))
	assert.Equal(t, 12, cart.GetTotal())
	cart = cart.Add(banana)
	assert.Equal(t, banana, cart.products[1])
	assert.Equal(t, 2, len(cart.products))
	assert.Equal(t, 30, cart.GetTotal())

}
