package logger

import "log/slog"

type Config struct {
	Level   slog.Level
	Console ConsoleConfig
	File    FileConfig
}

type ConsoleConfig struct {
	Enabled    bool
	TimeFormat string
}

type FileConfig struct {
	Enabled    bool
	Path       string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

func DefaultConfig() Config {
	return Config{
		Level: slog.LevelInfo,
		Console: ConsoleConfig{
			Enabled:    true,
			TimeFormat: "2006-01-02 15:04:05",
		},
	}
}
