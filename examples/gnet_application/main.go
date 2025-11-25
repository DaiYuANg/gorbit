package main

import (
	"log"
	"log/slog"
	"time"

	"github.com/DaiYuANg/gorbit"
	gorbit_gnet "github.com/DaiYuANg/gorbit/gnet"
	"github.com/DaiYuANg/gorbit/logger/zap_logger"
	"github.com/panjf2000/gnet/v2"
	"github.com/samber/oops"
)

type testServer struct {
	gnet.BuiltinEventEngine
	eng       gnet.Engine
	addr      string
	multicore bool
}

func (es *testServer) OnBoot(eng gnet.Engine) gnet.Action {
	es.eng = eng
	log.Printf("echo v_server with multi-core=%t is listening on %s\n", es.multicore, es.addr)
	return gnet.None
}

func (es *testServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)
	_, err := c.Write(buf)
	if err != nil {
		log.Println(oops.Wrap(err))
		return 0
	}
	log.Printf("conn %s read %s\n", c.RemoteAddr(), string(buf))
	return gnet.None
}

func (es *testServer) OnShutdown(eng gnet.Engine) {
	log.Println("echo v_server is closing")
}

func (es *testServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Printf("conn %s open\n", c.RemoteAddr())
	return nil, 0
}

func (es *testServer) OnTick() (delay time.Duration, action gnet.Action) {
	return 0, 0
}

func main() {
	container, err := gorbit.CreateContainerWithFxLogger(
		zap_logger.NewModule(),
		gorbit_gnet.NewModule(&testServer{}, slog.Default()),
	)
	if err != nil {
		panic(err)
	}
	container.Run()
}
