package config

type AppConfig struct {
	Mode Mode
	Name string `json:"name" koanf:"name"`
	Env  string `json:"env"  koanf:"env"`
}

func defaultConfig() *AppConfig {
	return &AppConfig{
		Mode: Dev,
	}
}
