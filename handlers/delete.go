package handlers

import (
	"Learning_Microservices/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// swagger:route DELETE /products/{id} products deleteProducts
// Return a list of products
// responses:
// 201: noContent

// DeleteProducts deletes a product from the database

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle DELETE Product", id)

	err := data.DeleteProduct(id)

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
	}
}
