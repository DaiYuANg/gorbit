package logger

import (
	"log/slog"

	"github.com/rs/zerolog"
	slogzerolog "github.com/samber/slog-zerolog/v2"
)

func newSlogHandler(zlogger zerolog.Logger, cfg Config) slog.Handler {
	return slogzerolog.Option{
		Level:     cfg.Level,
		Logger:    &zlogger,
		AddSource: true,
	}.NewZerologHandler()
}
