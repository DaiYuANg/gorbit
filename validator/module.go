package validator

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

var Module = fx.Module("validator", fx.Provide(newValidator))

func newValidator(option ...validator.Option) *validator.Validate {
	return validator.New(option...)
}
