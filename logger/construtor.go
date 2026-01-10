package logger

import "log/slog"

type Closer func() error

func NewLogger(opts ...Option) (*slog.Logger, Closer, error) {
	cfg := DefaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	zlogger, closeFn, err := newZerolog(cfg)
	if err != nil {
		return nil, nil, err
	}

	handler := newSlogHandler(zlogger, cfg)
	logger := slog.New(handler)

	return logger, closeFn, nil
}
