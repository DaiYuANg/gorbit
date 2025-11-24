package http

import (
	"fmt"

	"github.com/danielgtaylor/huma/v2"
)

// 功能型 Option
type HumaOption func(*humaConfig)

type humaConfig struct {
	Title    string
	Version  string
	DocsPath string
	Servers  []*huma.Server
}

// 默认配置
func defaultHumaConfig(title, version string, port int) *humaConfig {
	return &humaConfig{
		Title:    title,
		Version:  version,
		DocsPath: "/",
		Servers: []*huma.Server{
			{URL: fmt.Sprintf("http://localhost:%d", port)},
			{URL: fmt.Sprintf("http://127.0.0.1:%d", port)},
		},
	}
}

// Option helpers
func WithDocsPath(path string) HumaOption {
	return func(cfg *humaConfig) { cfg.DocsPath = path }
}

func WithServers(servers []*huma.Server) HumaOption {
	return func(cfg *humaConfig) { cfg.Servers = servers }
}

// 构建 Huma 配置
func NewHumaConfig(title, version string, port int, opts ...HumaOption) *huma.Config {
	cfg := defaultHumaConfig(title, version, port)
	for _, o := range opts {
		o(cfg)
	}

	humaCfg := huma.DefaultConfig(cfg.Title, cfg.Version)
	humaCfg.DocsPath = cfg.DocsPath
	humaCfg.Servers = cfg.Servers

	return &humaCfg
}
