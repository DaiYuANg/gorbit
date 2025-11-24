package bun

import "github.com/uptrace/bun/driver/sqliteshim"

// DatabaseOption 功能型 Option
type DatabaseOption func(*dbOptions)

type dbOptions struct {
	DriverName string
	DSN        string
	Debug      bool
}

// 默认值
func defaultDbOptions() *dbOptions {
	return &dbOptions{
		DriverName: sqliteshim.ShimName,
		DSN:        "file::memory:?cache=shared",
		Debug:      true,
	}
}

// Option helpers
func WithDriver(driver string) DatabaseOption {
	return func(o *dbOptions) { o.DriverName = driver }
}

func WithDSN(dsn string) DatabaseOption {
	return func(o *dbOptions) { o.DSN = dsn }
}

func WithDebug(debug bool) DatabaseOption {
	return func(o *dbOptions) { o.Debug = debug }
}
