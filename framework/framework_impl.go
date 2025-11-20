package framework

import (
	"fmt"
	"log/slog"

	"github.com/DaiYuANg/gorbit/pkg"
	_ "github.com/joho/godotenv/autoload"
	"github.com/samber/oops"
)

func (f *Framework) Run() error {
	for _, phase := range defaultPhases {
		if err := f.runPhase(phase); err != nil {
			return err
		}
	}
	return nil
}

func (f *Framework) Stop() error {
	f.ctx.Publish(EventFrameworkStop, nil)

	for _, m := range f.modules {
		err := m.Stop(f.ctx)
		if err != nil {
			return oops.Wrap(err)
		}
	}

	f.ctx.Publish(EventFrameworkStopDone, nil)
	return nil
}

func (f *Framework) runPhase(cfg PhaseConfig) error {
	appID := f.appID

	slog.Info(fmt.Sprintf("Framework Module %s Start", pkg.Capitalize(cfg.Phase.String())), "phase", cfg.Phase, "appID", appID)
	f.ctx.Publish(cfg.EventStart, nil)

	for _, m := range f.modules {
		slog.Info(fmt.Sprintf("%s module", cfg.Action[0]), "phase", cfg.Phase, "module", m.Name(), "appID", appID)
		if err := cfg.Handler(f, m); err != nil {
			slog.Error(fmt.Sprintf("%s module failed", cfg.Action[0]), "phase", cfg.Phase, "module", m.Name(), "error", err, "appID", appID)
			return fmt.Errorf("module %s %s failed: %w", m.Name(), cfg.Phase, err)
		}
		slog.Info(fmt.Sprintf("%s module", cfg.Action[1]), "phase", cfg.Phase, "module", m.Name(), "appID", appID)
	}

	f.ctx.Publish(cfg.EventDone, nil)
	slog.Info(fmt.Sprintf("Framework Module %s Done", pkg.Capitalize(cfg.Phase.String())), "phase", cfg.Phase, "appID", appID)

	return nil
}
