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
		ReadTimeout:    setting.Server.ReadTimeout * time.Second,
		WriteTimeout:   setting.Server.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
