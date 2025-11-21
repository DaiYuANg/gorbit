package zap_logger

import (
	"fmt"

	"github.com/DaiYuANg/gorbit/framework"
	"github.com/samber/do/v2"

	"log/slog"

	slogzap "github.com/samber/slog-zap/v2"
	"go.uber.org/zap"
)

type LoggerModule struct {
	logger *slog.Logger
	framework.Priority
}

func NewLoggerModule() *LoggerModule {
	return &LoggerModule{Priority: framework.NewPriority(framework.PriorityHigh)}
}

func (l *LoggerModule) Name() string {
	return "LoggerModule"
}

// Register 阶段：可在 injector 中注册 logger
func (l *LoggerModule) Register(i do.Injector) error {
	// 使用 zap 创建生产日志
	zapLogger, err := zap.NewProduction()
	if err != nil {
		return fmt.Errorf("failed to create zap logger: %w", err)
	}

	// 将 zap 转为 slog handler
	l.logger = slog.New(slogzap.Option{
		Level:  slog.LevelDebug,
		Logger: zapLogger,
	}.NewZapHandler())

	// 注册到 injector，这样其他模块可以注入 logger
	do.Provide[*slog.Logger](i, func(injector do.Injector) (*slog.Logger, error) {
		return l.logger, nil
	})
	return nil
}

// Init 阶段：可做一些 logger 配置
func (l *LoggerModule) Init(ctx *framework.AppContext) error {
	l.logger = l.logger.With("app", "example")
	l.logger.Info("LoggerModule initialized")
	return nil
}

// Start 阶段：可选
func (l *LoggerModule) Start(ctx *framework.AppContext) error {
	l.logger.Info("LoggerModule started")
	return nil
}

// Stop 阶段：可选
func (l *LoggerModule) Stop(ctx *framework.AppContext) error {
	l.logger.Info("LoggerModule stopped")
	return nil
}
