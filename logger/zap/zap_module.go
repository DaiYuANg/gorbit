package zap_logger

import (
	"os"

	"github.com/DaiYuANg/gorbit/framework"
	"github.com/samber/do/v2"
	"go.uber.org/zap/zapcore"

	"log/slog"

	slogzap "github.com/samber/slog-zap/v2"
	"go.uber.org/zap"
)

type LoggerModule struct {
	logger *slog.Logger
	framework.Prioritizer
}

func NewLoggerModule() *LoggerModule {
	return &LoggerModule{}
}

func (l *LoggerModule) Priority() int {
	return framework.NewPriority(framework.PriorityHigh).Value
}
func (l *LoggerModule) Name() string {
	return "LoggerModule"
}

// Register 阶段：可在 injector 中注册 logger
func (l *LoggerModule) Register(i do.Injector) error {
	// 使用 zap 创建生产日志
	zapLogger := newLogger()

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

func newLogger() *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig() // 可用 zap.NewProductionEncoderConfig() 视需求选择
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg), // 关键：使用 ConsoleEncoder 而不是 JSONEncoder
		zapcore.AddSync(os.Stdout),
		zapcore.InfoLevel,
	)

	logger := zap.New(core)
	return logger
}
