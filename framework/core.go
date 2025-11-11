package framework

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/samber/do/v2"
)

type Options struct {
	ConfigPath string
}

type Framework struct {
	Injector do.Injector
	Options  Options
}
