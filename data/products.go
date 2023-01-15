package data

import "time"

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Decription string  `json:"decription"`
	Price      float32 `json:"price"`
	SKU        string  `json:"sku"`
	CreatedOn  string  `json:"-"`
	UpdatedOn  string  `json:"-"`
	DeletedOn  string  `json:"-"`
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	&Product{
		ID:         1,
		Name:       "Latte",
		Decription: "Frothy milky coffee",
		Price:      2.45,
		SKU:        "abc123",
		CreatedOn:  time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
	},

	&Product{
		ID:         2,
		Name:       "Espresso",
		Decription: "Short and strong coffee without milk",
		Price:      2.00,
		SKU:        "cde456",
		CreatedOn:  time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
	},
}
