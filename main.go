package main

import (
	"context"
	"fmt"
	"gin-template/pkg/setting"
	"gin-template/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := routers.InitRouter()

	server := setting.Server

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", server.HttpPort),
		Handler:        router,
		ReadTimeout:    server.ReadTimeout * time.Second,
		WriteTimeout:   server.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	signalChan:= make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	// 主进程挂起，等待接受中断信号
	<-signalChan

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
