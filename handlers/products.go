package handlers

import (
	"Learning_Microservices/data"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// A list of products return in the response
// swagger:response productsResponse
type productResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}
type productsNoContent struct {
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	//The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

type Products struct {
	l *log.Logger
	v *data.Validation
}

func NewProduct(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

type KeyProduct struct{}

type ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /product/[id]")

type GenericError struct {
	Message string `json:"message"`
}
type ValidationError struct {
	Message []string `json:"message"`

}

func getProductID(r *http.Request) int {
	vars := mux.Vars(r)


	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	return id
}