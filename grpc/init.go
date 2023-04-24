package grpc

import (
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/niubir/area-service/config"
	"github.com/niubir/area-service/grpc/server"
	"github.com/niubir/area-service/logs"
	"github.com/niubir/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func Init(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	if !config.Config.GRPC.Enable {
		return
	}

	address := fmt.Sprintf(":%d", config.Config.GRPC.Port)

	setGRPCLogger()
	s := grpc.NewServer()
	server.RegisterAreaServiceServer(s, &AreaService{})
	listen, err := net.Listen("tcp", address)
	if err != nil {
		logs.System.Errorln("start grpc failed:", err)
		return
	}

	logs.System.Infof("start grpc, address(%s)\n", address)
	s.Serve(listen)
}

func setGRPCLogger() {
	grpcLogger := io.Discard
	if config.Config.GRPC.Debug {
		if w, err := logs.NewLogger("area_service_grpc"); err != nil {
			logs.System.Errorln("create area service grpc logger failed:", err)
		} else {
			grpcLogger = &logger.GRPCLogger{Logger: w}
		}
	}
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(grpcLogger, grpcLogger, grpcLogger))
}
