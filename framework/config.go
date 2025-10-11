package framework

import (
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config interface {
	String(key string) string
	Int(key string) int
	Bool(key string) bool
	Get(key string) any
	Unmarshal(out any) error
}

type KoanfConfig struct {
	k *koanf.Koanf
}

func (c *KoanfConfig) String(key string) string { return c.k.String(key) }
func (c *KoanfConfig) Int(key string) int       { return c.k.Int(key) }
func (c *KoanfConfig) Bool(key string) bool     { return c.k.Bool(key) }
func (c *KoanfConfig) Get(key string) any       { return c.k.Get(key) }
func (c *KoanfConfig) Unmarshal(out any) error  { return c.k.Unmarshal("", out) }

func NewKoanfLoader(configFile, envPrefix string) func() (Config, error) {
	return func() (Config, error) {
		k := koanf.New(".")

		// 1️⃣ 环境变量
		if envPrefix != "" {
			_ = k.Load(env.Provider(envPrefix+"_", ".", func(s string) string {
				return strings.ReplaceAll(strings.ToLower(
					strings.TrimPrefix(s, envPrefix+"_"),
				), "_", ".")
			}), nil)
		}

		// 2️⃣ 可选配置文件
		if configFile != "" {
			if err := k.Load(file.Provider(configFile), yaml.Parser()); err != nil {
				return nil, err
			}
		}

		return &KoanfConfig{k}, nil
	}
}
