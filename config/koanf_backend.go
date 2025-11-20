package config

import (
	"github.com/joho/godotenv"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/samber/oops"
)

type KoanfBackend struct {
	k *koanf.Koanf
}

func BackendKoanf() Backend {
	return &KoanfBackend{k: koanf.New(".")}
}

func (b *KoanfBackend) Load(opts Options) error {
	// dotenv
	if opts.EnvFile != "" {
		_ = godotenv.Load(opts.EnvFile)
	}

	// files
	if !opts.OnlyEnv {
		for _, f := range opts.Files {
			var p koanf.Parser

			switch getExt(f) {
			case ".yaml", ".yml":
				p = yaml.Parser()
			case ".json":
				p = json.Parser()
			case ".toml":
				p = toml.Parser()
			default:
				continue
			}

			if err := b.k.Load(file.Provider(f), p); err != nil {
				return oops.Wrap(err)
			}
		}
	}

	// env
	if err := b.k.Load(env.Provider(opts.EnvPrefix, ".", func(s string) string {

		return s
	}), nil); err != nil {
		return oops.Wrap(err)
	}

	return nil
}

func (b *KoanfBackend) Unmarshal(path string, target any) error {
	return b.k.Unmarshal(path, target)
}

func (b *KoanfBackend) Raw() any {
	return b.k
}

func getExt(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			return path[i:]
		}
	}
	return ""
}
