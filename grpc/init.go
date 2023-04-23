package grpc

import (
	"fmt"
	"net"
	"sync"

	"github.com/niubir/area-service/config"
	"github.com/niubir/area-service/grpc/server"
	"github.com/niubir/area-service/logs"
	"google.golang.org/grpc"
)

func Init(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	if !config.Config.GRPC.Enable {
		return
	}

	address := fmt.Sprintf(":%d", config.Config.GRPC.Port)

	s := grpc.NewServer()
	server.RegisterAreaServiceServer(s, &AreaService{})
	listen, err := net.Listen("tcp", address)
	if err != nil {
		logs.Error("start grpc failed: %v", err)
		return
	}

	logs.Info("start grpc, address(%s)", address)
	s.Serve(listen)
}
