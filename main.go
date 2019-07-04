package main

import (
	"fmt"
	"gin-template/pkg/setting"
	"gin-template/routers"
	"net/http"
	"time"
)

var db = make(map[string]string)

func main() {
	router := routers.InitRouter()

	// 配置
	s := http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Server.HttpPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
