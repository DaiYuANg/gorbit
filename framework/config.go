package framework

type AppConfig struct {
	Mode Mode
}

func defaultConfig() *AppConfig {
	return &AppConfig{
		Mode: Dev,
	}
}
