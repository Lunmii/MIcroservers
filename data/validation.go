package data

import (
	"fmt"
	"github.com/go-playground/validator"
	"regexp"
)

type ValidatorError struct {
	validator.FieldError
}

func (v ValidatorError) Error() string {
	return fmt.Sprintf(
		"Key: %s' Error: Field validation for %s' failed on the %s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

type ValidationErrors []ValidatorError

func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return &Validation{validate}

}

// Validate the item
// for more detail the returned error can be cast into a
// validator.ValidationErrors collection
//
// if ve, ok := err.(validator.ValidationErrors); ok {
//			fmt.Println(ve.Namespace())
//			fmt.Println(ve.Field())
//			fmt.Println(ve.StructNamespace())
//			fmt.Println(ve.StructField())
//			fmt.Println(ve.Tag())
//			fmt.Println(ve.ActualTag())
//			fmt.Println(ve.Kind())
//			fmt.Println(ve.Type())
//			fmt.Println(ve.Value())
//			fmt.Println(ve.Param())
//			fmt.Println()
//	}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs := v.validate.Struct(i).(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	var returnErrs []ValidationError
	for _, err := range errs {
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}

func validateSKU(f1 validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	sku := re.FindAllString(f1.Field().String(), -1)

	if len(sku) == 1 {
		return true
	}

	return false
}
