package config

import flag "github.com/spf13/pflag"

// 功能型 Option
type Option[T any] func(*configOptions[T])

type configOptions[T any] struct {
	Default   T
	JSONFiles []string
	YAMLFiles []string
	TOMLFiles []string
	EnvPrefix string
	FlagSets  []*flag.FlagSet
}

// 默认值
func defaultConfigOptions[T any](defaultStruct T) *configOptions[T] {
	return &configOptions[T]{
		Default:   defaultStruct,
		JSONFiles: []string{},
		YAMLFiles: []string{},
		TOMLFiles: []string{},
		EnvPrefix: "APP_",
		FlagSets:  []*flag.FlagSet{},
	}
}
