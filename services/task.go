package services

import (
	"time"

	"github.com/niubir/area-service/config"
	"github.com/niubir/area-service/logs"
)

func initTask() error {
	go autoFreshArea()
	return nil
}

func autoFreshArea() {
	if config.Config.AutoFreshTime == "" {
		return
	}
	autoFreshTime, err := time.Parse("15:04:05", config.Config.AutoFreshTime)
	if err != nil {
		return
	}
	now := time.Now()
	next := time.Date(
		now.Year(), now.Month(), now.Day(),
		autoFreshTime.Hour(), autoFreshTime.Minute(), autoFreshTime.Second(), autoFreshTime.Nanosecond(),
		now.Location(),
	)
	if now.After(next) {
		next = next.AddDate(0, 0, 1)
	}

	timer := time.NewTimer(next.Sub(now))
	for range timer.C {
		go func() {
			if err := FreshAreas(); err != nil {
				logs.System.Errorln("auto fresh areas failed:", err)
			}
		}()
		timer.Reset(24 * time.Hour)
	}
}
