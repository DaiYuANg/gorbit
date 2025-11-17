package config

type Options struct {
	Backend   Backend    // KoanfBackend or ViperBackend
	Files     []string   // config.yaml, config.json...
	EnvFile   string     // .env
	EnvPrefix string     // APP_
	OnlyEnv   bool       // true = no file load
	App       *AppConfig // core config
}
