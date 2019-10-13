package main

import "errors"

type Product struct {
	price int
	name  string
}

func NewProduct(name string, price int) (Product, error) {
	if price <= 0 && name == "" {
		return Product{}, errors.New("a product's price needs to be a positive non-zero value and it needs to have a name")
	}
	if price <= 0 {
		return Product{}, errors.New("a product's price cannot be zero or negative")
	}
	if name == "" {
		return Product{}, errors.New("a product needs to have a name")
	}
	return Product{
		name: name,
		price: price}, nil
}

type Cart struct {
	products []Product
	total int
}

func NewCart() Cart {
	return Cart{}
}

func (cart Cart) Add(product Product) Cart {
	cart.products = append(cart.products, product)
	cart.total = cart.total + product.price
	return cart
}

func (cart Cart) GetTotal() int {
	return cart.total
}
