package framework

func New[Config any](c Config) *Framework[Config] {
	return &Framework[Config]{
		userCfg: c,
	}
}
