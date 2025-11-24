package gnet

// 配置
type Config struct {
	Addr      string
	Multicore bool
}

type Option func(*Config)

func WithAddr(addr string) Option {
	return func(c *Config) { c.Addr = addr }
}

func WithMulticore(multicore bool) Option {
	return func(c *Config) { c.Multicore = multicore }
}
