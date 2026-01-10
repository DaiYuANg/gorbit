package config

import flag "github.com/spf13/pflag"

// 功能型 Option helper
func WithJSONSupport[T any](files ...string) Option[T] {
	return func(o *configOptions[T]) { o.JSONFiles = files }
}

func WithYAMLSupport[T any](files ...string) Option[T] {
	return func(o *configOptions[T]) { o.YAMLFiles = files }
}

func WithTOMLSupport[T any](files ...string) Option[T] {
	return func(o *configOptions[T]) { o.TOMLFiles = files }
}

func WithEnvPrefix[T any](prefix string) Option[T] {
	return func(o *configOptions[T]) { o.EnvPrefix = prefix }
}

func WithFlagSet[T any](fs *flag.FlagSet) Option[T] {
	return func(o *configOptions[T]) {
		o.FlagSets = append(o.FlagSets, fs)
	}
}
