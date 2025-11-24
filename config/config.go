package config

// 功能型 Option
type ConfigOption[T any] func(*configOptions[T])

type configOptions[T any] struct {
	Default   T
	JSONFiles []string
	YAMLFiles []string
	TOMLFiles []string
	EnvPrefix string
}

// 默认值
func defaultConfigOptions[T any](defaultStruct T) *configOptions[T] {
	return &configOptions[T]{
		Default:   defaultStruct,
		JSONFiles: []string{"mock/mock.json"},
		YAMLFiles: []string{"mock/mock.yml"},
		TOMLFiles: []string{"mock/mock.toml"},
		EnvPrefix: "APP_",
	}
}
