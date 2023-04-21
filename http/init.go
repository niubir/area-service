package http

import (
	"io"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/niubir/area-service/config"
	"github.com/niubir/area-service/logs"
)

func Init(wg *sync.WaitGroup) {
	setGinMode()
	setGinWriter()

	engine := gin.Default()

	initRoutes(engine)

	engine.Run(":10011")

	wg.Done()
}

func setGinWriter() {
	ginWriter := io.Discard
	ginErrorWriter := io.Discard
	if config.Config.Debug {
		if w, err := logs.NewFileLogger("area_service_http"); err != nil {
			logs.Error("create area service http logger failed: %v", err)
		} else {
			ginWriter = w
		}
		if w, err := logs.NewFileLogger("area_service_http_error"); err != nil {
			logs.Error("create area service error http logger failed: %v", err)
		} else {
			ginErrorWriter = w
		}
	}
	gin.DefaultWriter = ginWriter
	gin.DefaultErrorWriter = ginErrorWriter
}

func setGinMode() {
	var ginMode string
	switch {
	case config.Config.IsProduction():
		ginMode = gin.ReleaseMode
	case config.Config.IsTest():
		ginMode = gin.TestMode
	default:
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)
}
