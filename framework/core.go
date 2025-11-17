package framework

import (
	"github.com/go-co-op/gocron/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/samber/do/v2"
	goeventbus "github.com/stanipetrosyan/go-eventbus"
)

type Options struct {
	ConfigPath string
}

type Framework struct {
	injector  do.Injector
	options   Options
	modules   []Module
	scheduler *gocron.Scheduler
	eventBus  *goeventbus.EventBus
}
