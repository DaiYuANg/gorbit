package schedule

import (
	"context"
	"log/slog"

	"github.com/go-co-op/gocron/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("scheduleModule", fx.Provide(newScheduler), fx.Invoke(startScheduler))

func newScheduler(logger *slog.Logger) (gocron.Scheduler, error) {
	gocronLogger := NewGocronSlogLogger(logger)
	return gocron.NewScheduler(
		gocron.WithLogger(gocronLogger),
	)
}

func startScheduler(lc fx.Lifecycle, cron gocron.Scheduler, logger *slog.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go cron.Start()
			logger.Info("Startup scheduler")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return cron.Shutdown()
		},
	})
}
