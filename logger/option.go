package logger

import "log/slog"

type Option func(*Config)

func WithLevel(level slog.Level) Option {
	return func(c *Config) {
		c.Level = level
	}
}

func WithConsole(enabled bool) Option {
	return func(c *Config) {
		c.Console.Enabled = enabled
	}
}

func WithFile(cfg FileConfig) Option {
	return func(c *Config) {
		c.File = cfg
		c.File.Enabled = true
	}
}
