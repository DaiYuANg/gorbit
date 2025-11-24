package bun

import (
	"database/sql"
	"fmt"

	"github.com/samber/oops"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/extra/bundebug"
)

// NewDatabase 创建 bun.DB
func NewDatabase(opts ...DatabaseOption) *bun.DB {
	options := defaultDbOptions()
	for _, o := range opts {
		o(options)
	}

	sqldb, err := sql.Open(options.DriverName, options.DSN)
	if err != nil {
		panic(
			oops.Wrap(
				fmt.Errorf("failed to open database: %w", err),
			),
		)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	if options.Debug {
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
			bundebug.FromEnv("BUNDEBUG"),
		))
	}

	return db
}
