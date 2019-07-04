package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	// 获得Gin的实例Engine
	r := gin.Default()

	// Context是Gin最重要的部分
	r.GET("/ping", func(c *gin.Context) {
		// String将字符串写入响应体
		c.String(http.StatusOK, "ping %s", "pong")
	})

	return r
}
