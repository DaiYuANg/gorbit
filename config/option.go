package config

type ConfigOptions struct {
	Path         string // 配置文件路径，可选
	Format       string // 配置文件格式：yaml|json|toml|env 等
	EnvPrefix    string // 环境变量前缀，可选
	EnvDelimiter string // 环境变量分隔符，如 "_" 或 "__"
}
