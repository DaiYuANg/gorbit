package validator

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

func NewValidatorModule(option ...validator.Option) fx.Option {
	return fx.Module("validator", fx.Provide(
		func() *validator.Validate {
			return validator.New(option...)
		},
	))
}
