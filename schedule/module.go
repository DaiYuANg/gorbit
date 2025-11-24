package schedule

import (
	"context"

	"github.com/go-co-op/gocron/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("scheduleModule", fx.Provide(newScheduler), fx.Invoke(startScheduler))

func newScheduler() (gocron.Scheduler, error) {
	return gocron.NewScheduler(
		gocron.WithLogger(
			gocron.NewLogger(gocron.LogLevelInfo),
		),
	)
}

func startScheduler(lc fx.Lifecycle, cron gocron.Scheduler) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go cron.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return cron.Shutdown()
		},
	})
}
