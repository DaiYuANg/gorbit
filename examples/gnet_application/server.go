package main

import (
	"log/slog"
	"time"

	"github.com/panjf2000/gnet/v2"
	"github.com/samber/oops"
)

type testServer struct {
	gnet.BuiltinEventEngine
	eng       gnet.Engine
	addr      string
	multicore bool
	logger    *slog.Logger
}

func (es *testServer) OnBoot(eng gnet.Engine) gnet.Action {
	es.eng = eng
	es.logger.Info("echo v_server with multi-core=%t is listening on", es.multicore, es.addr)
	return gnet.None
}

func (es *testServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)
	_, err := c.Write(buf)
	if err != nil {
		es.logger.Error("Error", oops.Wrap(err))
		return 0
	}
	es.logger.Info("conn %s read", c.RemoteAddr(), string(buf))
	return gnet.None
}

func (es *testServer) OnShutdown(eng gnet.Engine) {
	es.logger.Info("echo test_server is closing")
}

func (es *testServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	es.logger.Info("conn %s open", c.RemoteAddr())
	return nil, 0
}

func (es *testServer) OnTick() (delay time.Duration, action gnet.Action) {
	return 0, 0
}
