package main

import (
	"sync"

	"github.com/niubir/area-service/http"

	"github.com/niubir/area-service/config"
	"github.com/niubir/area-service/logs"
	"github.com/niubir/area-service/services"
)

func main() {
	Init()

	var wg sync.WaitGroup

	wg.Add(1)
	go http.Init(&wg)

	wg.Wait()
}

func Init() {
	if err := config.Init(); err != nil {
		panic(err)
	}
	if err := logs.Init(); err != nil {
		panic(err)
	}
	if err := services.Init(); err != nil {
		panic(err)
	}
}
