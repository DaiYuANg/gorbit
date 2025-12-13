package zap_logger

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"syscall"

	"github.com/samber/oops"
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

	level := zap.NewAtomicLevelAt(zapcore.Level(cfg.Level))

	consoleCore := zapcore.NewCore(
		consoleEncoder,
		zapcore.AddSync(os.Stdout),
		level,
	)

	fileCore := zapcore.NewCore(
		fileEncoder,
		zapcore.AddSync(lumberJackLogger),
		level,
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

// deferLogger 安全注册 logger 的 OnStop Hook
// 支持 stdout/stderr 和文件日志，忽略常见无效文件描述符错误
func deferLogger(lc fx.Lifecycle, logger *zap.Logger) {
	lc.Append(
		fx.Hook{
			OnStop: func(ctx context.Context) error {
				// 封装 sync 函数
				safeSync := func(w io.Writer) error {
					if f, ok := w.(*os.File); ok {
						if err := f.Sync(); err != nil &&
							!errors.Is(err, syscall.EBADF) &&
							!errors.Is(err, syscall.EINVAL) {
							return err
						}
					}
					return nil
				}

				// 尝试同步 stdout/stderr
				if err := safeSync(os.Stdout); err != nil {
					return oops.With("stdout sync failed").Wrap(err)
				}
				if err := safeSync(os.Stderr); err != nil {
					return oops.With("stderr sync failed").Wrap(err)
				}

				// 尝试同步 zap 内部的 fileCore（如果是文件日志）
				if err := logger.Sync(); err != nil &&
					!errors.Is(err, syscall.EBADF) &&
					!errors.Is(err, syscall.EINVAL) {
					return oops.With("logger file sync failed").Wrap(err)
				}

				return nil
			},
		},
	)
}
