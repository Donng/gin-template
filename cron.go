package main

import (
	"gin-template/models"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	// 每分钟执行一次
	c.AddFunc("0 * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})

	c.Start()

	t1 := time.NewTimer(10 * time.Second)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
