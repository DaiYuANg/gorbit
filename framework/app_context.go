package framework

import (
	"github.com/samber/do/v2"
	goeventbus "github.com/stanipetrosyan/go-eventbus"
)

type AppContext struct {
	injector do.Injector
	events   goeventbus.EventBus
}

func newAppContext(i do.Injector, e goeventbus.EventBus) *AppContext {
	return &AppContext{injector: i, events: e}
}

// Publish 事件发布
func (c *AppContext) Publish(topic string, payload any) {
	message := goeventbus.NewMessageBuilder().SetPayload(payload).Build()
	c.events.Channel(topic).Publisher().Publish(message)
}

// On 事件订阅
func (c *AppContext) On(topic string, handler func(ctx goeventbus.Context)) {
	c.events.Channel(topic).Subscriber().Listen(handler)
}

func (c *AppContext) Injector() do.Injector {
	return c.injector
}
