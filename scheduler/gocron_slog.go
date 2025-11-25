package schedule

import (
	"log/slog"

	"github.com/go-co-op/gocron/v2"
)

type GocronSlogLogger struct {
	internalLogger *slog.Logger
}

func (l *GocronSlogLogger) Debug(msg string, args ...any) {
	l.internalLogger.Debug(msg, args...)
}
func (l *GocronSlogLogger) Error(msg string, args ...any) {
	l.internalLogger.Error(msg, args...)
}
func (l *GocronSlogLogger) Info(msg string, args ...any) {
	l.internalLogger.Info(msg, args...)
}
func (l *GocronSlogLogger) Warn(msg string, args ...any) {
	l.internalLogger.Warn(msg, args...)
}

func NewGocronSlogLogger(log *slog.Logger) gocron.Logger {
	return &GocronSlogLogger{
		internalLogger: log,
	}
}
