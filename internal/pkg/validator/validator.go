package validator

import "github.com/go-playground/validator/v10"

var v = validator.New()

func Default() *validator.Validate {
	return v
}

func RequiredParams(params ...interface{}) error {
	for _, param := range params {
		err := v.Var(param, "required")
		if err != nil {
			return err
		}
	}
	return nil
}
