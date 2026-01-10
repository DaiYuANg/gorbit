package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/samber/lo"
	oopszerolog "github.com/samber/oops/loggers/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func newZerolog(cfg Config) (zerolog.Logger, Closer, error) {
	var writers []io.Writer
	var closers []io.Closer

	if cfg.Console.Enabled {
		writers = append(writers, zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: cfg.Console.TimeFormat,
		})
	}

	if cfg.File.Enabled {
		lj := &lumberjack.Logger{
			Filename:   cfg.File.Path,
			MaxSize:    cfg.File.MaxSize,
			MaxAge:     cfg.File.MaxAge,
			MaxBackups: cfg.File.MaxBackups,
		}
		writers = append(writers, lj)
		closers = append(closers, lj)
	}

	if len(writers) == 0 {
		writers = append(writers, os.Stdout)
	}

	level, err := zerolog.ParseLevel(cfg.Level.String())
	if err != nil {
		level = zerolog.InfoLevel
	}

	logger := zerolog.New(io.MultiWriter(writers...)).
		Level(level).
		With().
		Timestamp().
		Logger()

	zerolog.ErrorStackMarshaler = oopszerolog.OopsStackMarshaller
	zerolog.ErrorMarshalFunc = oopszerolog.OopsMarshalFunc

	return logger, func() error {
		lo.ForEach(closers, func(item io.Closer, _ int) {
			_ = item.Close()
		})
		return nil
	}, nil
}
