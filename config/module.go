package config

import (
	"log/slog"
	"strings"

	flag "github.com/spf13/pflag"

	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/toml/v2"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"go.uber.org/fx"
)

func NewConfigModule[T any](defaultStruct T, opts ...Option[T]) fx.Option {
	options := defaultConfigOptions(defaultStruct)
	lo.ForEach(opts, func(o Option[T], _ int) {
		o(options)
	})

	return fx.Module("config_module",
		fx.Provide(func() *koanf.Koanf {
			return koanf.New(".")
		}),
		fx.Provide(func(k *koanf.Koanf, logger *slog.Logger) (*T, error) {
			def := options.Default

			// 先加载默认 struct
			if err := k.Load(structs.Provider(def, "koanf"), nil); err != nil {
				return nil, oops.Wrap(err)
			}

			// 加载 JSON 文件
			lo.ForEach(options.JSONFiles, func(f string, _ int) {
				if err := k.Load(file.Provider(f), json.Parser()); err != nil {
					logger.Warn("error loading JSON config :",
						slog.String("key", f),
						slog.Any("config load error", oops.Wrap(err)))
				}
			})

			lo.ForEach(options.YAMLFiles, func(f string, _ int) {
				if err := k.Load(file.Provider(f), yaml.Parser()); err != nil {
					logger.Warn("error loading YAML config: ",
						slog.String("key", f),
						slog.String("config load error", oops.Wrap(err).Error()),
					)
				}
			})

			lo.ForEach(options.TOMLFiles, func(f string, _ int) {
				if err := k.Load(file.Provider(f), toml.Parser()); err != nil {
					logger.Warn("error loading TOML config %s: %v",
						slog.String("key", f),
						slog.Any("config load error", oops.Wrap(err)),
					)
				}
			})

			// 加载环境变量
			mapEnvKey := func(s string) string {
				return strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(s, options.EnvPrefix)), "_", ".")
			}
			if err := k.Load(env.Provider(options.EnvPrefix, ".", mapEnvKey), nil); err != nil {
				return nil, err
			}
			// 6️⃣ 命令行参数
			if len(options.FlagSets) > 0 {
				lo.ForEach(options.FlagSets, func(fs *flag.FlagSet, _ int) {
					if err := k.Load(posflag.Provider(fs, ".", nil), nil); err != nil {
						logger.Warn("error loading CLI flags",
							slog.Any("config load error", oops.Wrap(err)),
						)
					}
				})
			}
			// 映射到结构体
			if err := k.Unmarshal("", &def); err != nil {
				return nil, err
			}

			logger.Info("loaded config:", slog.Any("RAW", k.All()))
			logger.Info("loaded config:", slog.Any("Struct", def))
			return &def, nil
		}),
	)
}
