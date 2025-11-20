package http

import "github.com/samber/do/v2"

type Module struct {
	Server Server
	Addr   string
}

func (m *Module) Init(i do.Injector) error {
	println("xxx")
	return nil
}

func (m *Module) Name() string { return "http" }

func (m *Module) Register(i do.Injector) error {
	// 注册 HTTP Server
	do.Provide[Server](i, func(injector do.Injector) (Server, error) {
		return m.Server, nil
	})
	return nil
}

func (m *Module) Start(i do.Injector) error {
	s, err := do.Invoke[Server](i)
	if err != nil {
		return err
	}
	return s.Listen(m.Addr)
}

func (m *Module) Stop(i do.Injector) error {
	s, err := do.Invoke[Server](i)
	if err != nil {
		return err
	}
	return s.Shutdown()
}
