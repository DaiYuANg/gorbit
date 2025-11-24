package config

import (
	"strings"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// NewConfigModule 返回 fx.Module
// NewConfigModule 泛型实现
func NewConfigModule[T any](defaultStruct T, opts ...ConfigOption[T]) fx.Option {
	options := defaultConfigOptions(defaultStruct)
	for _, o := range opts {
		o(options)
	}

	return fx.Module("config_module",
		fx.Provide(func() *koanf.Koanf {
			return koanf.New(".")
		}),
		fx.Provide(func(k *koanf.Koanf, logger *zap.SugaredLogger) (*T, error) {
			def := options.Default

			// 先加载默认 struct
			if err := k.Load(structs.Provider(def, "koanf"), nil); err != nil {
				return nil, err
			}

			// 加载 JSON 文件
			for _, f := range options.JSONFiles {
				if err := k.Load(file.Provider(f), json.Parser()); err != nil {
					logger.Warnf("error loading JSON config %s: %v", f, err)
				}
			}

			// 加载 YAML 文件
			for _, f := range options.YAMLFiles {
				if err := k.Load(file.Provider(f), yaml.Parser()); err != nil {
					logger.Warnf("error loading YAML config %s: %v", f, err)
				}
			}

			// 加载 TOML 文件
			for _, f := range options.TOMLFiles {
				if err := k.Load(file.Provider(f), toml.Parser()); err != nil {
					logger.Warnf("error loading TOML config %s: %v", f, err)
				}
			}

			// 加载环境变量
			mapEnvKey := func(s string) string {
				return strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(s, options.EnvPrefix)), "_", ".")
			}
			if err := k.Load(env.Provider(options.EnvPrefix, ".", mapEnvKey), nil); err != nil {
				return nil, err
			}

			// 映射到结构体
			if err := k.Unmarshal("", &def); err != nil {
				return nil, err
			}

			logger.Infof("loaded config: %+v", def)
			return &def, nil
		}),
	)
}
