package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/do/v2"
	"github.com/spf13/viper"
)

type ViperProvider struct {
	v *viper.Viper
}

func (v *ViperProvider) Get(path string) any {
	return v.v.Get(path)
}
func (v *ViperProvider) String(path string) string {
	return v.v.GetString(path)
}
func (v *ViperProvider) Int(path string) int {
	return v.v.GetInt(path)
}
func (v *ViperProvider) Bool(path string) bool {
	return v.v.GetBool(path)
}
func (v *ViperProvider) Unmarshal(path string, out any) error {
	return v.v.UnmarshalKey(path, out)
}
func (v *ViperProvider) Source() string {
	return "viper"
}

// NewViper creates a new ConfigProvider using spf13/viper.
func NewViper(opts ...ConfigOptions) func(i do.Injector) error {
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

		v := viper.New()

		// === 文件加载 ===
		v.SetConfigFile(o.Path)
		v.SetConfigType(o.Format)

		if err := v.ReadInConfig(); err != nil {
			// 文件不存在时允许继续（例如纯环境变量配置）
			if !os.IsNotExist(err) {
				return err
			}
		}

		// === 环境变量加载 ===
		// Viper 内部用 "_" 作为分隔符，可通过替换符实现兼容
		v.SetEnvPrefix(o.EnvPrefix)
		v.SetEnvKeyReplacer(strings.NewReplacer(o.EnvDelimiter, "_"))
		v.AutomaticEnv()

		provider := &ViperProvider{v: v}
		do.ProvideValue[ConfigProvider](i, provider)
		do.ProvideValue(i, v) // 也注入原生 *viper.Viper
		return nil
	}
}
