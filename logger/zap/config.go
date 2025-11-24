package zap_logger

import "go.uber.org/zap/zapcore"

type Config struct {
	// log file path
	FilePath string

	// log level
	Level zapcore.Level

	// rotate configs
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool

	// encoders
	ConsoleJSON bool
	FileJSON    bool
}

func defaultConfig() *Config {
	return &Config{
		FilePath:   "",
		Level:      zapcore.DebugLevel,
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     7,
		Compress:   true,

		ConsoleJSON: false,
		FileJSON:    true,
	}
}

type Option func(*Config)

func WithFilePath(p string) Option {
	return func(c *Config) {
		c.FilePath = p
	}
}

func WithLevel(l zapcore.Level) Option {
	return func(c *Config) {
		c.Level = l
	}
}

func WithConsoleJSON() Option {
	return func(c *Config) {
		c.ConsoleJSON = true
	}
}

func WithFileJSON() Option {
	return func(c *Config) {
		c.FileJSON = true
	}
}

func WithRotate(maxSize, maxBackup, maxAge int) Option {
	return func(c *Config) {
		c.MaxSize = maxSize
		c.MaxBackups = maxBackup
		c.MaxAge = maxAge
	}
}

func WithCompress(b bool) Option {
	return func(c *Config) {
		c.Compress = b
	}
}
