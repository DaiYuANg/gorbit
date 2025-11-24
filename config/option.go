package config

// 功能型 Option helper
func WithJSONSupport[T any](files ...string) ConfigOption[T] {
	return func(o *configOptions[T]) { o.JSONFiles = files }
}

func WithYAMLSupport[T any](files ...string) ConfigOption[T] {
	return func(o *configOptions[T]) { o.YAMLFiles = files }
}

func WithTOMLSupport[T any](files ...string) ConfigOption[T] {
	return func(o *configOptions[T]) { o.TOMLFiles = files }
}

func WithEnvPrefix[T any](prefix string) ConfigOption[T] {
	return func(o *configOptions[T]) { o.EnvPrefix = prefix }
}
