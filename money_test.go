package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNewMoney(t *testing.T)  {
	money, err := NewMoney(44)
	require.NoError(t, err)
	assert.Equal(t, 44, money.Amount)
	money2, err := NewMoney(490)
	require.NoError(t, err)
	assert.Equal(t, 490, money2.Amount)
	_, err = NewMoney(-43)
	require.Error(t, err)
}

func TestAddMoney(t *testing.T)  {
	money1, _ := NewMoney(44)
	money2, _ := NewMoney(490)
	money3 := money1.Add(money2)
	assert.Equal(t, 534, money3.Amount)
}
