// Package handlers Package classification of Product API
//
// # Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// -application/json
//
// Produces:
// -application/json
// swagger:meta
package handlers

import "Learning_Microservices/data"

type errorResponseWrapper struct {
	Body GenericError
}

type errorValidationWrapper struct {
	Body ValidationError
}

type productsResponseWrapper struct {
	Body []data.Product
}
type noContentResponseWrapper struct {
}

type productParamsWrapper struct {
	Body data.Product
}

type productIDParamsWrapper struct {
	ID int `json:"id"`
}
