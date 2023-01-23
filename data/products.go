package data

import (
	"encoding/json"
	"fmt"
	"io"
)

// Product defines the structure for an API product
// swagger:model
var ErrProductNotFound = fmt.Errorf("Product not found")

type Product struct {
	// the id for this user
	//
	// required: true
	// min: 1

	ID         int     `json:"id"`
	Name       string  `json:"name" validate:"required"`
	Decription string  `json:"decription" `
	Price      float32 `json:"price" validate:"gt=0"`
	SKU        string  `json:"sku" validate:"required,sku"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func GetProductByID(id int) (*Product, error) {
	i := findIndexByProductId(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}

	return productList[i], nil
}

func AddProduct(p Product) {
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, &p)
}

func UpdateProduct(p Product) error {
	i := findIndexByProductId(p.ID)
	if i == -1 {
		return ErrProductNotFound
	}
	productList[i] = &p

	return nil
}

func DeleteProduct(id int) error {
	i := findIndexByProductId(id)
	if i == -1 {
		return ErrProductNotFound
	}
	productList = append(productList[:i], productList[i+1])

	return nil
}

func findIndexByProductId(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

var productList = []*Product{
	&Product{
		ID:         1,
		Name:       "Latte",
		Decription: "Frothy milky coffee",
		Price:      2.45,
		SKU:        "abc123",
	},

	&Product{
		ID:         2,
		Name:       "Espresso",
		Decription: "Short and strong coffee without milk",
		Price:      2.00,
		SKU:        "cde456",
	},
}
