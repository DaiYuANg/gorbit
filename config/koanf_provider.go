package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/samber/do/v2"
)

type KoanfProvider struct {
	k *koanf.Koanf
}

func (k *KoanfProvider) Get(path string) any {
	return k.k.Get(path)
}
func (k *KoanfProvider) String(path string) string {
	return k.k.String(path)
}
func (k *KoanfProvider) Int(path string) int {
	return k.k.Int(path)
}
func (k *KoanfProvider) Bool(path string) bool {
	return k.k.Bool(path)
}
func (k *KoanfProvider) Unmarshal(path string, out any) error {
	return k.k.Unmarshal(path, out)
}
func (k *KoanfProvider) Source() string {
	return "koanf"
}

// 自动选择解析器
func parserFor(format string) koanf.Parser {
	switch strings.ToLower(format) {
	case "yaml", "yml":
		return yaml.Parser()
	case "json":
		return json.Parser()
	case "toml":
		return toml.Parser()
	default:
		return yaml.Parser()
	}
}

// NewKoanf 创建一个可配置的 Koanf 模块
func NewKoanf(opts ...ConfigOptions) func(i do.Injector) error {
	return func(i do.Injector) error {
		var o ConfigOptions
		if len(opts) > 0 {
			o = opts[0]
		}

		// === 默认值 ===
		if o.Path == "" {
			if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
				o.Path = envPath
			} else {
				o.Path = "config.yaml"
			}
		}
		if o.Format == "" {
			o.Format = strings.TrimPrefix(filepath.Ext(o.Path), ".")
			if o.Format == "" {
				o.Format = "yaml"
			}
		}
		if o.EnvDelimiter == "" {
			o.EnvDelimiter = "_"
		}

		k := koanf.New(".")

		// 从文件加载（可选）
		if err := k.Load(file.Provider(o.Path), parserFor(o.Format)); err != nil {
			// 允许文件不存在，使用空配置
			if !os.IsNotExist(err) {
				return err
			}
		}

		// 从环境变量加载（覆盖文件配置）
		prefix := o.EnvPrefix
		_ = k.Load(env.Provider(prefix, o.EnvDelimiter, func(s string) string {
			// ENV 变量自动转小写
			return strings.ToLower(strings.TrimPrefix(s, prefix))
		}), nil)

		provider := &KoanfProvider{k: k}

		do.ProvideValue[ConfigProvider](i, provider)
		do.ProvideValue(i, k) // 同时注入原始 koanf
		return nil
	}
}
