package config

import (
	"github.com/joho/godotenv"
	"github.com/samber/oops"
	"github.com/spf13/viper"
)

type ViperBackend struct {
	v *viper.Viper
}

func BackendViper() Backend {
	v := viper.New()
	return &ViperBackend{v: v}
}

func (b *ViperBackend) Load(opts Options) error {
	// dotenv
	if opts.EnvFile != "" {
		_ = godotenv.Load(opts.EnvFile)
	}

	b.v.SetEnvPrefix(opts.EnvPrefix)
	b.v.AutomaticEnv()

	// Files
	if !opts.OnlyEnv {
		for _, f := range opts.Files {
			b.v.SetConfigFile(f)
			if err := b.v.MergeInConfig(); err != nil {
				return oops.Wrap(err)
			}
		}
	}
	return nil
}

func (b *ViperBackend) Unmarshal(path string, target any) error {
	return b.v.Unmarshal(target)
}

func (b *ViperBackend) Raw() any {
	return b.v
}
