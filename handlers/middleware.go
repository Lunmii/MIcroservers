package handlers

import (
	"Learning_Microservices/data"
	"context"
	"net/http"
)

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		errs := p.v.Validate(prod)
		if len(errs) != 0 {

			p.l.Println("[ERROR] validating product", errs)

			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Message: err.Error()}, rw)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)

	})
}
