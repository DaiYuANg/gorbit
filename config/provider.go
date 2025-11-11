package config

// ConfigProvider 是框架内所有配置解析器的统一接口
type ConfigProvider interface {
	Get(path string) any
	String(path string) string
	Int(path string) int
	Bool(path string) bool
	Unmarshal(path string, out any) error
	Source() string
}
