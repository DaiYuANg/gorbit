package framework

import (
	"github.com/daiyuang/gorbit/config"
	"github.com/daiyuang/gorbit/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	EnvPrefix string
}

func New[Config any](c Config, options ...Options) *Framework[Config] {
	return &Framework[Config]{
		app: fx.New(
			config.Module,
			logger.Module,
			fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
				fxLogger := &fxevent.ZapLogger{Logger: log}
				fxLogger.UseLogLevel(zapcore.DebugLevel)
				return fxLogger
			}),
		),
	}
}
