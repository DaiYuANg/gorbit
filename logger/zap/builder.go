package zap_logger

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func newLogger(cfg *Config) *zap.Logger {
	logFile := cfg.FilePath
	if logFile == "" {
		logFile = filepath.Join(os.TempDir(), "warden.log")
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}

	// console encoder
	consoleEncoderCfg := zap.NewProductionEncoderConfig()
	consoleEncoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	consoleEncoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoderCfg.TimeKey = "T"
	consoleEncoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var consoleEncoder zapcore.Encoder
	if cfg.ConsoleJSON {
		consoleEncoder = zapcore.NewJSONEncoder(consoleEncoderCfg)
	} else {
		consoleEncoder = zapcore.NewConsoleEncoder(consoleEncoderCfg)
	}

	// file encoder
	fileEncoderCfg := zap.NewProductionEncoderConfig()
	fileEncoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	fileEncoderCfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	fileEncoderCfg.TimeKey = "timestamp"
	fileEncoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var fileEncoder zapcore.Encoder
	if cfg.FileJSON {
		fileEncoder = zapcore.NewJSONEncoder(fileEncoderCfg)
	} else {
		fileEncoder = zapcore.NewConsoleEncoder(fileEncoderCfg)
	}

	consoleCore := zapcore.NewCore(
		consoleEncoder,
		zapcore.AddSync(os.Stdout),
		cfg.Level,
	)

	fileCore := zapcore.NewCore(
		fileEncoder,
		zapcore.AddSync(lumberJackLogger),
		cfg.Level,
	)

	core := zapcore.NewTee(consoleCore, fileCore)

	return zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)
}

func sugaredLogger(log *zap.Logger) *zap.SugaredLogger {
	return log.Sugar()
}

func deferLogger(lc fx.Lifecycle, logger *zap.Logger) {
	lc.Append(
		fx.Hook{
			OnStop: func(context.Context) error {
				if err := logger.Sync(); err != nil && !errors.Is(err, syscall.EINVAL) {
					return fmt.Errorf("logger sync failed: %v", err)
				}
				return nil
			},
		},
	)
}
