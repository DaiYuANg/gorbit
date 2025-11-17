package http

type Server interface {
	Listen(addr string) error
	Shutdown() error
}
