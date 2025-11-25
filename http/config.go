package http

import "strconv"

type Config struct {
	Port int `koanf:"port"`
}

type APIDoc struct {
	Enable bool   `koanf:"enable"`
	Path   string `koanf:"path"`
}

func (h Config) GetPort() string {
	return strconv.Itoa(h.Port)
}
