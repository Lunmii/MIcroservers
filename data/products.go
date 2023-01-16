package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

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

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
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
