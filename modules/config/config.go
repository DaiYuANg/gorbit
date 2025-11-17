package config

import "go.uber.org/zap/zapcore"

type AppConfig struct {
	Mode     Mode
	LogLevel zapcore.Level
}

func defaultConfig() *AppConfig {
	return &AppConfig{
		Mode: Dev,
	}
}
