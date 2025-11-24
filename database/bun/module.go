package bun

import (
	"github.com/uptrace/bun"
	"go.uber.org/fx"
)

func NewDatabaseModule(opts ...DatabaseOption) fx.Option {
	return fx.Module("database_module",
		fx.Provide(func() *bun.DB {
			return NewDatabase(opts...)
		}),
	)
}
