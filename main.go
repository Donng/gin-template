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

	server := setting.Server

	s := http.Server{
		Addr:           fmt.Sprintf(":%d", server.HttpPort),
		Handler:        router,
		ReadTimeout:    server.ReadTimeout * time.Second,
		WriteTimeout:   server.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
